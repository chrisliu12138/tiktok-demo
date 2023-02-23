package controller

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/middleware/ffmpeg"
	"SimpleDouyin/middleware/ftp"
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
	data, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	if err != nil {
		panic("获取用户名失败")
	}
	//根据token获取userName
	userId := c.GetString("userId")
	Id, err := strconv.ParseInt(userId, 10, 64)
	user, err := dao.GetTableUserById(Id)
	if err != nil {
		panic("根据userName查询用户失败")
	}
	//获取token的用户
	finalName := fmt.Sprintf("%d_%s", user.Id, filename) //存的文件名字bear.mp4
	//savFile public\1_bear1.mp4   finalName:1_bear1.mp4
	var saveFile = filepath.Join("./public/", finalName) //文件路径
	//存到本地
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//将本地视频上传到服务器
	ftp.Ftp(saveFile, finalName)
	//截图
	ffmpeg.Ffmpeg(finalName, finalName+"-output")
	//给video表添加一条记录，包括title  playUrl uerId等
	impl := service.VideoServiceImpl{}
	var flag = impl.Add(user.Id, "/ftpfile/"+filename, "test")
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

	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
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
