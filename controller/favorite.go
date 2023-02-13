package controller

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid

func FavoriteAction(c *gin.Context) {
	userid := c.Query("userId")
	vedioId := c.Query("vedioId")

	ip := c.ClientIP()
	service.LimitIP(ip, vedioId)

	var like = 0
	if userid != "" && vedioId != "" {
		like = service.Like(userid, vedioId)
	}
	if like == 1 {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "点赞失败"})
	}
}

//
//// FavoriteList all users have same favorite video list
//func FavoriteList(c *gin.Context) {
//
//	userid := c.Query("userId")
//
//	Videoslist := service.GetVedioLikeList(userid)
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: dao.Response{
//			StatusCode: 0,
//		},
//		VideoList: Videoslist,
//	})
//}
