package service

import (
	"SimpleDouyin/config"
	"SimpleDouyin/dao"
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*
返回当前用户的所有点赞视频列表
*/
func GetVedioLikeList(userId string) []dao.Video {
	members := dao.SMembers(userId)
	IdList := make([]int64, len(members))
	if members == nil {
		return nil
	}
	for i, s := range members {
		IdList[i], _ = strconv.ParseInt(s, 10, 64)
	}

	//调用service方法
	impl := VideoServiceImpl{}
	videos := impl.QueryListByVedioIdList(IdList)
	return videos
}

/*
返回当前用户的获赞总数，不存在则返回0
*/
func GetUserLikeCount(userId string) int64 {
	id, _ := strconv.ParseInt(userId, 10, 64)
	impl := VideoServiceImpl{}
	videos := impl.Query(id)
	fmt.Println(videos)
	count := int64(0)
	for _, s := range videos {
		id := strconv.Itoa(int(s.Id))
		likeCount := GetVedioLikeCount(id)
		fmt.Println(likeCount)
		count += likeCount
	}
	return count
}

/*
返回当前视频的点赞总数，视频不存在则返回0set GOOS=linux
*/
func GetVedioLikeCount(vedioId string) int64 {
	count := dao.SelectCount(vedioId)
	return count
}

/*
*
点赞操作
返回值为int，点赞成功则为1，否则为0且报错
*/
func Like(vedioId, userId string) int64 {
	add := dao.SAddVideoLike(vedioId, userId)
	add2 := dao.SAddUserLike(userId, vedioId)
	return add & add2
}

/*
*
取消点赞操作
返回值为int，点赞成功则为1，否则为0且报错
*/
func DislikeVedio(vedioId, userId string) int64 {
	remove := dao.SremoveVedioLike(vedioId, userId)
	sremove := dao.SremoveUserLike(userId, vedioId)
	return remove & sremove
}

// 查询当前用户是否点赞
func LikeVedioOrNot(vedioId, userId string) bool {
	member := dao.SIsMember(vedioId, userId)
	return member
}

// 定时任务，把redis中的数据更新到mysql中的点赞表里
func TimeMission() {
	ticker := time.NewTicker(config.UPDATE_PERIOD)
	go func() {
		for {
			<-ticker.C
			SaveRedisDataToMySql()
		}
	}()
}

func SaveRedisDataToMySql() {
	count := dao.GetVedioCount()
	var i int64
	for i = 0; i < count; i += config.MYSQL_LIMIT {
		//拿到当前的所有视频id
		userIds := dao.GetVedioIdWithLimit(i, config.MYSQL_LIMIT)
		//根据视频id更新mysql表中所有视频的点赞数
		for _, id := range userIds {
			//1. 从redis中取出当前vedio的点赞数
			likeCount := GetVedioLikeCount(string(id))
			//2. 更新mysql中的对应id的vedio点赞数
			bool := dao.UpdateVedioLikeCount(id, likeCount)
			if !bool {
				errors.New("redis持久化更新失败")
			}
		}
	}
}

/*
测试用，为redis添加数据
*/
func Add(vedioId, userId string) int64 {
	result := dao.SAddVideoLike(vedioId, userId)
	return result
}

/*
测试用，为redis添加数据
*/
func AdduserId(userId, vedioId string) int64 {
	add := dao.SAddUserLike(userId, vedioId)
	return add
}
