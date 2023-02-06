package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	dao.Response
	VideoList [30]dao.Video `json:"video_list,omitempty"`
	NextTime  time.Time     `json:"next_time,omitempty"`
}

// GET  /douyin/feed
func Feed(c *gin.Context) {
	//1.查询前30个video
	rows := impl.QueryAll()
	if rows == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: dao.Response{StatusCode: 0, StatusMsg: "查询失败"},
		})
	}
	var length int = len(rows)
	//2.查询成功则封装Response
	var videoList [30]dao.Video
	for i := 0; i < length; i++ {
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
	nextTime := rows[len(rows)-1].CreatedAt
	//3.把封装的结果返回给前端
	c.JSON(http.StatusOK, FeedResponse{
		Response:  dao.Response{StatusCode: 1, StatusMsg: "查询成功"},
		VideoList: videoList,
		NextTime:  nextTime, //取结果的最后一个
	})
}
