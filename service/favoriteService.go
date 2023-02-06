package service

import (
	"container/list"
	"context"
	"github.com/RaymondCode/simple-demo/Utils"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/controller"
)

/*
返回当前用户的所有点赞视频列表
*/
var Vedio_like = config.Vedio_like
var User_like = config.User_like

var RDB = Utils.RDB

func GetVedioLikeList(userId string) []controller.Video {
	resultList := list.List{}
	ctx := context.Background()
	result, err := RDB.SMembers(ctx, User_like+userId).Result()
	if err != nil {
		panic(err)
	}
	for _, s := range result {
		resultList.PushFront(s)
	}
	//调用service方法
	return nil
}

/*
返回当前视频的点赞总数，视频不存在则返回0
*/
func GetVedioLikeCount(vedioId string) int64 {
	var count = int64(0)
	ctx := context.Background()
	result, err := RDB.SCard(ctx, Vedio_like+vedioId).Result()
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
	result, err := RDB.SAdd(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	result2, err := RDB.SAdd(ctx, User_like+userId, vedioId).Result()
	if err != nil {
		panic(err)
	}
	return int(result) & int(result2)
}
func Add(vedioId, userId string) int64 {
	ctx := context.Background()
	result, err := RDB.SAdd(ctx, Vedio_like+vedioId, userId).Result()
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
	result, err := RDB.SRem(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	result2, err := RDB.SRem(ctx, User_like+userId, vedioId).Result()
	if err != nil {
		panic(err)
	}
	return int(result) & int(result2)
}

//查询当前用户是否点赞
func LikeVedioOrNot(vedioId, userId string) bool {
	ctx := context.Background()
	result, err := RDB.SIsMember(ctx, Vedio_like+vedioId, userId).Result()
	if err != nil {
		panic(err)
	}
	return result
}
