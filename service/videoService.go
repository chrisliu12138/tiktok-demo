package service

import "github.com/RaymondCode/simple-demo/dao"

type VideoService interface {
	/*
		API
	*/
	// 查询最新的30个稿件
	QueryAll() [30]dao.Video

	// 根据videoArray查询稿件
	QueryListByVedionl(videoArray []int64) []dao.Video

	// 根据userID查询稿件
	Query(userid int64) []dao.Video

	// 上传稿件   userID:是谁发的(根据token去用户信息表查id)  playUrl：存在哪里  title：视频名字是啥
	Add(userId int64, playUrl string, title string) bool
}
