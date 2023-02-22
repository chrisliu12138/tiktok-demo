package service

import (
	"SimpleDouyin/config"
	"SimpleDouyin/dao"
	"errors"
	"strconv"
	"sync"
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
返回当前视频的点赞总数，视频不存在则返回0
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

// TotalFavourite 根据userId获取这个用户总共被点赞数量
func TotalFavourite(userId int64) (int64, error) {
	//根据userId获取这个用户的发布视频列表信息
	videoIdList := Query(userId)
	var sum int64 //该用户的总被点赞数
	//提前开辟空间,存取每个视频的点赞数
	videoLikeCountList := new([]int64)
	//采用协程并发将对应videoId的点赞数添加到集合中去
	i := len(videoIdList)
	var wg sync.WaitGroup
	wg.Add(i)
	wg.Wait()
	//遍历累加，求总被点赞数
	for _, count := range *videoLikeCountList {
		sum += count
	}
	return sum, nil
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
