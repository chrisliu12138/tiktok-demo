package controller

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/service"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	dao.Response
	VideoList []dao.Video `json:"video_list"`
}

// POST /douyin/publish/action
func Publish(c *gin.Context) {
	//是否传过来视频数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	//获取登录用户

	userName := c.PostForm("token")
	if err != nil {
		panic("获取用户名失败")
	}

	user, err := dao.GetTableUserByUserName(userName)
	if err != nil {
		panic("根据userName查询用户失败")
	}
	//获取token的用户
	finalName := fmt.Sprintf("%d_%s", user.Id, filename) //存的文件名字
	var saveFile = filepath.Join("./public/", finalName) //文件路径
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//给video表添加一条记录，包括title  playUrl uerId等
	impl := service.VideoServiceImpl{}
	var flag = impl.Add(user.Id, saveFile, "test")
	fmt.Println(flag)

	if flag == true {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded successfully",
		})
	} else {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded failed",
		})
	}

}

// GET /douyin/publish/list
func PublishList(c *gin.Context) {
	//1.获取登录用户

	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		panic("字符型转整型失败")
	}

	user, err := dao.GetTableUserById(userId)
	if err != nil {
		panic("根据userId查询用户失败")
	}
	//2.根据用户id查询其所有Video
	impl := service.VideoServiceImpl{}
	videoList := impl.Query(int64(user.Id))
	if videoList == nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: dao.Response{StatusCode: 0, StatusMsg: "查询失败"},
		})
	}

	//3.结果返回给前端
	c.JSON(http.StatusOK, FeedResponse{
		Response:  dao.Response{StatusCode: 1, StatusMsg: "查询成功"},
		VideoList: videoList,
	})

}
