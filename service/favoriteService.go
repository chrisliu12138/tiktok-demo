package service

import (
	"context"
	"github.com/RaymondCode/simple-demo/Utils"
	"github.com/RaymondCode/simple-demo/service/impl"
	"strconv"

	"github.com/RaymondCode/simple-demo/config"
)

var Vedio_like = config.Vedio_like
var User_like = config.User_like

/*
返回当前用户的所有点赞视频列表
*/
func GetVedioLikeList(userId string) []impl.Result {
	var IdList []int64
	ctx := context.Background()
	result, err := Utils.RDB.SMembers(ctx, User_like+userId).Result()
	if err != nil {
		panic(err)
	}
	for i, s := range result {
		IdList[i], err = strconv.ParseInt(s, 10, 64)
	}
	//调用service方法
	//vedioList := []dao.Video
	results := impl.QueryListByVedionl(IdList)
	//UuserServiceImpl := UserServiceImpl{}
	//for _, r := range results {
	//	vedio := dao.Video{
	//		Id: int64(r.ID),
	//		Author:
	//	}
	//
	//}
	return results
}

//
func LimitIP(ip, vedioId string) bool {
	ctx := context.Background()
	Utils.RDB.SCard(ctx, Vedio_like+vedioId).Result()
	return true
}

/*
返回当前视频的点赞总数，视频不存在则返回0
*/
func GetVedioLikeCount(vedioId string) int64 {
	var count = int64(0)
	ctx := context.Background()
	result, err := Utils.RDB.SCard(ctx, Vedio_like+vedioId).Result()
	if err != nil {
		panic(err)
	}
	count = result
	return count
}

/**
点赞操作
返回值为int，点赞成功则为1，否则为0且报错
*/
func Like(vedioId, userId string) int {
	ctx := context.Background()
	result, err := Utils.RDB.SAdd(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	result2, err := Utils.RDB.SAdd(ctx, User_like+userId, vedioId).Result()
	if err != nil {
		panic(err)
	}
	return int(result) & int(result2)
}

/*
测试用，为redis添加数据
*/
func Add(vedioId, userId string) int64 {
	ctx := context.Background()
	result, err := Utils.RDB.SAdd(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	return result
}

/*
测试用，为redis添加数据
*/
func AdduserId(userId, vedioId string) int64 {
	ctx := context.Background()
	result, err := Utils.RDB.SAdd(ctx, User_like+userId, vedioId).Result()
	if err != nil {
		panic(err)
	}
	return result
}

/**
取消点赞操作
返回值为int，点赞成功则为1，否则为0且报错
*/
func DislikeVedio(vedioId, userId string) int {
	ctx := context.Background()
	result, err := Utils.RDB.SRem(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	result2, err := Utils.RDB.SRem(ctx, User_like+userId, vedioId).Result()
	if err != nil {
		panic(err)
	}
	return int(result) & int(result2)
}

//查询当前用户是否点赞
func LikeVedioOrNot(vedioId, userId string) bool {
	ctx := context.Background()
	result, err := Utils.RDB.SIsMember(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	return result
}
