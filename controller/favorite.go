package controller

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid

func FavoriteAction(c *gin.Context) {
	userid := c.Query("userId")
	vedioId := c.Query("vedioId")
	action_type := c.Query("action_type")

	//ip限流，查询当前ip是否超过设置的访问次数
	ip := c.ClientIP()
	islimitIP := dao.LimitIP(ip, vedioId)
	if !islimitIP {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 2, StatusMsg: "点赞失败，操作太频繁"})
	}

	like := 0
	if userid != "" && vedioId != "" {
		if action_type == "1" {
			like = int(service.Like(userid, vedioId))
		} else {
			like = int(service.DislikeVedio(userid, vedioId))
		}
	}
	if like == 1 {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "点赞失败"})
	}
}

// FavoriteList 返回用户的所有点赞过的视频
func FavoriteList(c *gin.Context) {
	userid := c.Query("userId")

	Videoslist := service.GetVedioLikeList(userid)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
		VideoList: Videoslist,
	})
}
