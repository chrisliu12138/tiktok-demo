package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	dao.Response
	VideoList []dao.Video `json:"video_list"`
}

// demo鉴权
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      1,
	},
}

// POST /douyin/publish/action
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	//鉴权  默认存在 token =zhangleidouyin
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
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
	user := usersLoginInfo[token]                        //获取token的用户
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
	//1.鉴权

	//a.鉴权失败
	token := c.PostForm("token")
	//鉴权  默认存在 token =zhangleidouyin
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: dao.Response{
				StatusCode: 0,
				StatusMsg:  "请先登录",
			},
		})
		return
	}
	//鉴权成功往后走
	user := usersLoginInfo[token] //获取token的用户
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
