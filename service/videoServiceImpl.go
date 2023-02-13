package service

import (
	"SimpleDouyin/dao"
	"fmt"
)

type VideoServiceImpl struct {
}

// 上传稿件   userID:是谁发的(根据token去用户信息表查id)  playUrl：存在哪里  title：视频名字是啥
func (VideoServiceImpl *VideoServiceImpl) Add(userId int64, playUrl string, title string) bool {
	flag := dao.Add(uint(userId), playUrl, title)
	return flag
}

// 根据userID查询稿件
func (VideoServiceImpl *VideoServiceImpl) Query(userid int64) []dao.Video {
	//1.根据用户id查询其所有Video
	var videoList []dao.Video
	rows := dao.Query(uint(userid))
	if rows != nil {
		//2.查询成功则封装Response
		videoList = convertPOtoDTO(rows, len(rows))
		return videoList
	}
	return nil

}

// 根据videoArray查询稿件
func (VideoServiceImpl *VideoServiceImpl) QueryListByVedioIdList(videoArray []int64) []dao.Video {
	var videoList []dao.Video
	rows := dao.QueryListByVedionl(videoArray)
	if rows != nil {
		videoList = convertPOtoDTO(rows, len(rows))
		return videoList
	}
	return nil

}

// 查询最新的30个稿件
func (VideoServiceImpl *VideoServiceImpl) QueryAll() []dao.Video {
	var videoList []dao.Video
	//封装videoList
	rows := dao.QueryAll()
	if rows != nil {
		videoList = convertPOtoDTO(rows, 30)
		return videoList
	}

	return nil
}

// 转换PO为DTO
func convertPOtoDTO(rows []dao.Result, len int) []dao.Video {
	var videoList []dao.Video
	for i := 0; i < len; i++ {
		fmt.Println(rows[i])
		var author dao.User
		author = dao.User{
			Id:            int64(rows[i].UserID),
			Name:          rows[i].Name,
			FollowCount:   int64(rows[i].FollowCount),
			FollowerCount: int64(rows[i].FollowerCount),
			IsFollow:      int(rows[i].IsFollow),
		}
		//封装author
		videoList[i].Author = author
		videoList[i].Id = int64(rows[i].ID)
		videoList[i].CoverUrl = rows[i].CoverUrl
		videoList[i].CommentCount = int64(rows[i].CommentCount)
		videoList[i].FavoriteCount = int64(rows[i].FavoriteCount)
		videoList[i].IsFavorite = int(rows[i].IsFavorite)
		videoList[i].PlayUrl = rows[i].PlayUrl
		videoList[i].Title = rows[i].Title
	}
	return videoList
}
