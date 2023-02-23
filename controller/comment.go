package controller

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	dao.Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	dao.Response
	Comment dao.Comment `json:"comment,omitempty"`
}

// CommentAction 新增/删除评论
func CommentAction(c *gin.Context) {
	//从上下文中获取执行当前操作的用户的id
	id := c.Query("user_id")
	userid, _ := strconv.ParseInt(id, 10, 64)

	//获取执行操作的视频id
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "获取视频id失败"})
		log.Println("出现无法解析成64位整数的视频id")
		return
	}

	//获取当前操作类型
	actionType := c.Query("action_type")

	if actionType == "1" { //增加评论
		commentText := c.Query("comment_text")
		comment, err := service.Comment(videoId, userid, commentText)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				dao.Response{StatusCode: 1, StatusMsg: "评论异常"},
				dao.Comment{},
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				dao.Response{StatusCode: 0, StatusMsg: "评论成功"},
				comment,
			})
			return
		}
	} else if actionType == "2" { //删除评论
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "获取评论id失败"})
			log.Println("出现无法解析成64位整数的视频id")
			return
		}
		err = service.DeleteComment(commentId)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				dao.Response{StatusCode: 1, StatusMsg: "评论异常"},
				dao.Comment{},
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				dao.Response{StatusCode: 0, StatusMsg: "success"},
				dao.Comment{},
			})
		}
	}

}

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	//确认视频id无误
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "获取视频id失败"})
		log.Println("出现无法解析成64位整数的视频id")
		return
	}

	//获取commentList
	commentList, err := service.GetCommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: dao.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			CommentList: nil,
		})
	}

	//返回commentList
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    dao.Response{StatusCode: 0},
		CommentList: commentList,
	})
}
