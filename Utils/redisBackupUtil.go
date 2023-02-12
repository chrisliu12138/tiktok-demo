package Utils

//
//import (
//	"errors"
//	"github.com/RaymondCode/simple-demo/config"
//	"github.com/RaymondCode/simple-demo/dao"
//	"github.com/RaymondCode/simple-demo/service"
//	"time"
//)
//
//// 定时任务，把redis中的数据更新到mysql中的点赞表里
//func TimeMission() {
//	ticker := time.NewTicker(config.UPDATE_PERIOD)
//	go func() {
//		for {
//			<-ticker.C
//			SaveRedisDataToMySql()
//		}
//	}()
//}
//
//func SaveRedisDataToMySql() {
//	count := dao.GetVedioCount()
//	var i int64
//	for i = 0; i < count; i += config.MYSQL_LIMIT {
//		//拿到当前的所有视频id
//		userIds := dao.GetVedioIdWithLimit(i, config.MYSQL_LIMIT)
//		//根据视频id更新mysql表中所有视频的点赞数
//		for _, id := range userIds {
//			//1. 从redis中取出当前vedio的点赞数
//			likeCount := service.GetVedioLikeCount(string(id))
//			//2. 更新mysql中的对应id的vedio点赞数
//			bool := dao.UpdateVedioLikeCount(id, likeCount)
//			if !bool {
//				errors.New("redis持久化更新失败")
//			}
//		}
//	}
//}
