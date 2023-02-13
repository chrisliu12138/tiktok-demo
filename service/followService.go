package service

import (
	"SimpleDouyin/middleware/DBUtils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"SimpleDouyin/config"
	"SimpleDouyin/dao"
	"SimpleDouyin/middleware/rabbitmq"
	"github.com/jinzhu/gorm"
)

func Follow(userId, toUserid int64, actionType int64) (bool, error) {
	var result bool
	var err error
	switch {
	case actionType == 1:
		result, err = addFollow(userId, toUserid)
	case actionType == 2:
		result, err = unfollow(userId, toUserid)
	}
	return result, err
}

// 关注列表
func FollowList(userId int64) ([]dao.User, error) {
	//查询redis
	userIdStr := strconv.FormatInt(userId, 10)
	n, err := DBUtils.RdbFollow.SCard(DBUtils.Ctx, userIdStr).Result()
	if err != nil {
		return nil, err
	}
	//redis中存在记录
	if n > 0 {
		followUsers, err := DBUtils.RdbFollow.SMembers(DBUtils.Ctx, userIdStr).Result()
		if err != nil {
			return nil, err
		}
		return queryList(userId, followUsers, true)
	}
	toUsers := make([]int, 0)
	//从mysql中获取
	if err := DBUtils.Db.Raw("select to_user_id from follow where user_id = ? and cancel = 0", userIdStr).Scan(&toUsers).Error; err != nil {
		return nil, err
	}
	go addFollowToRedis(userId, toUsers)
	toUsersStr := make([]string, len(toUsers))
	for index, toUserId := range toUsers {
		toUsersStr[index] = strconv.Itoa(toUserId)
	}
	return queryList(userId, toUsersStr, true)
}

// 粉丝列表
func FollowerList(userId int64) ([]dao.User, error) {
	//查询redis
	userIdStr := strconv.FormatInt(userId, 10)
	n, err := DBUtils.RdbFollower.Exists(DBUtils.Ctx, userIdStr).Result()
	if err != nil {
		return nil, err
	}
	//redis中存在记录
	if n > 0 {
		followerUsers, err := DBUtils.RdbFollower.SMembers(DBUtils.Ctx, userIdStr).Result()
		if err != nil {
			return nil, err
		}
		return queryList(userId, followerUsers, false)
	}
	users := make([]int, 0)
	//从mysql中获取
	if err := DBUtils.Db.Raw("select user_id from follow where to_user_id = ?", userIdStr).Scan(&users).Error; err != nil {
		return nil, err
	}
	go addFollowerToRedis(userId, users)
	usersStr := make([]string, len(users))
	for index, userId := range users {
		usersStr[index] = strconv.Itoa(userId)
	}
	return queryList(userId, usersStr, false)
}

// follow为true代表请求关注列表，否则代表请求粉丝列表
func queryList(userId int64, list []string, follow bool) ([]dao.User, error) {
	users := make([]dao.User, 0)
	for _, user := range list {

		var tmpUser dao.User
		if err := DBUtils.Db.Raw("select id, name, follow_count, follower_count from users where id = ?", user).Scan(&tmpUser).Error; err != nil {
			return nil, err
		}

		if follow {
			if err := DBUtils.Db.Raw("select user_id from follow where user_id = ? and to_user_id = ? and cancel = 0", userId, user).Error; err == nil {
				tmpUser.IsFollow = 1
			} else if gorm.IsRecordNotFoundError(err) {
				log.Println("mysql查询错误: ", err)
			}
		} else {
			if err := DBUtils.Db.Raw("select user_id from follow where user_id = ? and to_user_id = ? and cancel = 0", user, userId).Error; gorm.IsRecordNotFoundError(err) {
				tmpUser.IsFollow = 1
			}
		}
		users = append(users, tmpUser)
	}
	return users, nil
}

func addFollow(userId, toUserId int64) (bool, error) {
	//先更新mysql，mysql插入成功后再将redis更新的消息放入rabbitmq，防止mysql出现回滚后，redis中出现脏数据

	userIdStr := strconv.FormatInt(userId, 10)
	toUserIdStr := strconv.FormatInt(toUserId, 10)

	sql := fmt.Sprintf("CALL followAction(%v,%v)", userIdStr, toUserIdStr)
	log.Printf("执行关注操作, SQL如下: %s", sql)
	if err := DBUtils.Db.Raw(sql).Scan(nil).Error; nil != err {
		log.Println(err.Error())
		return false, err
	}

	//向mq发送消息
	msg := strings.Builder{}
	msg.WriteString(userIdStr)
	msg.WriteString(" ")
	msg.WriteString(toUserIdStr)
	rabbitmq.FollowRmq.Publish(msg.String())

	return true, nil
}

func unfollow(userId, toUserId int64) (bool, error) {
	//redis更新
	//1、user的关注列表-1
	//2、target_user的粉丝列表-1
	userIdStr := strconv.FormatInt(userId, 10)
	toUserIdStr := strconv.FormatInt(toUserId, 10)

	sql := fmt.Sprintf("CALL unfollowAction(%v,%v)", userId, toUserId)
	log.Printf("执行取关操作, SQL如下: %s", sql)
	if err := DBUtils.Db.Raw(sql).Scan(nil).Error; nil != err {
		log.Println(err.Error())
		return false, err
	}

	//向mq发送消息
	msg := strings.Builder{}
	msg.WriteString(userIdStr)
	msg.WriteString(" ")
	msg.WriteString(toUserIdStr)
	rabbitmq.UnFollowRmq.Publish(msg.String())

	return true, nil
}

// 获取关注总数
func GetFollowCnt(userId int64) (followCnt int64, err error) {
	//查询redis
	userIdStr := strconv.FormatInt(userId, 10)
	followCnt, err = DBUtils.RdbFollow.SCard(DBUtils.Ctx, userIdStr).Result()
	if err != nil {
		log.Println("查询redis出错", err)
		return followCnt, err
	}
	//redis中存在记录
	if followCnt > 0 {
		return followCnt, nil
	}
	toUsers := make([]int, 0)
	//从mysql中获取
	if err := DBUtils.Db.Raw("select to_user_id from follow where user_id = ? and cancel = 0", userIdStr).Scan(&toUsers).Error; err != nil {
		log.Println("查询mysql出错", err)
		return followCnt, err
	}
	go addFollowToRedis(userId, toUsers)
	return int64(len(toUsers)), err
}

// 获取粉丝总数
func GetFollowerCnt(userId int64) (followCnt int64, err error) {
	//查询redis
	userIdStr := strconv.FormatInt(userId, 10)
	followCnt, err = DBUtils.RdbFollow.SCard(DBUtils.Ctx, userIdStr).Result()
	if err != nil {
		log.Println("查询redis出错", err)
		return followCnt, err
	}
	//redis中存在记录
	if followCnt > 0 {
		return followCnt, nil
	}
	toUsers := make([]int, 0)
	//从mysql中获取
	if err := DBUtils.Db.Raw("select user_id from follow where to_user_id = ? and cancel = 0", userIdStr).Scan(&toUsers).Error; err != nil {
		log.Println("查询mysql出错", err)
		return followCnt, err
	}
	go addFollowerToRedis(userId, toUsers)
	return int64(len(toUsers)), err
}

// 将关注列表放入redis
func addFollowToRedis(userId int64, toUsers []int) {
	userIdStr := strconv.FormatInt(userId, 10)
	for i := 0; i < len(toUsers); i++ {
		toUserIdStr := strconv.Itoa(toUsers[i])
		DBUtils.RdbFollow.SRem(DBUtils.Ctx, userIdStr, toUserIdStr)
		DBUtils.RdbFollow.Expire(DBUtils.Ctx, userIdStr, config.ExpireTime)
	}
}

// 将粉丝列表放入redis
func addFollowerToRedis(userId int64, toUsers []int) {
	userIdStr := strconv.FormatInt(userId, 10)
	for i := 0; i < len(toUsers); i++ {
		toUserIdStr := strconv.Itoa(toUsers[i])
		DBUtils.RdbFollower.SRem(DBUtils.Ctx, toUserIdStr, userIdStr)
		DBUtils.RdbFollower.Expire(DBUtils.Ctx, toUserIdStr, config.ExpireTime)
	}
}
