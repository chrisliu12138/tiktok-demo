package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
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
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		panic("字符型转整型失败")
	}

	user, err := dao.GetTableUserById(userId)
	if err != nil {
		panic("根据userId查询用户失败")
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
	var flag bool = impl.Add(uint(user.Id), saveFile, "test")
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
	//获取登录用户
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		panic("字符型转整型失败")
	}

	user, err := dao.GetTableUserById(userId)
	if err != nil {
		panic("根据userId查询用户失败")
	}
	//2.根据用户id查询其所有Video
	rows := impl.Query(uint(user.Id))
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
	//3.把封装的结果返回给前端
	c.JSON(http.StatusOK, FeedResponse{
		Response:  dao.Response{StatusCode: 1, StatusMsg: "查询成功"},
		VideoList: videoList,
	})

}
