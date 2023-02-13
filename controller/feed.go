package controller

import (
	impl2 "SimpleDouyin/dao"
	"SimpleDouyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	impl2.Response
	VideoList []impl2.Video `json:"video_list,omitempty"`
	//NextTime  time.Time     `json:"next_time,omitempty"`
}

// GET  /douyin/feed
func Feed(c *gin.Context) {
	//1.查询前30个video
	impl := service.VideoServiceImpl{}
	videoList := impl.QueryAll()
	if videoList == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: impl2.Response{StatusCode: 0, StatusMsg: "查询失败"},
		})
	}
	//2.返回给前端
	c.JSON(http.StatusOK, FeedResponse{
		Response:  impl2.Response{StatusCode: 1, StatusMsg: "查询成功"},
		VideoList: videoList,
	})
}
