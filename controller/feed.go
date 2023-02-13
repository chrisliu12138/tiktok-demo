package controller

import (
	impl2 "SimpleDouyin/dao"
	"SimpleDouyin/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	impl2.Response
	VideoList []impl2.Video `json:"video_list,omitempty"`
	NextTime  time.Time     `json:"next_time,omitempty"`
}

// GET  /douyin/feed
func Feed(c *gin.Context) {
	//1.查询前30个video
	videoList := service.VideoServiceImpl.QueryAll()
	if videoList == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: impl2.Response{StatusCode: 0, StatusMsg: "查询失败"},
		})
	}
	var length int = len(videoList)
	//2.返回给前端
	c.JSON(http.StatusOK, FeedResponse{
		Response:  impl2.Response{StatusCode: 1, StatusMsg: "查询成功"},
		VideoList: videoList,
		NextTime:  nextTime, //取结果的最后一个
	})
}
