package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type Response struct{
	StatusCode int32           `json:"status_code"`
	StatusMsg  string    			 `json:"status_msg"`
	UserList   []dao.User   `json:"user_list,omitempty"`
}

func FollowController(context *gin.Context) {
	var err error
	var userId int64
	if userId, err = strconv.ParseInt(context.GetString("userId"), 10, 64); err != nil{
		context.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  "user_id格式错误",
		})
		return
	}
	var toUserId int64
	if toUserId, err = strconv.ParseInt(context.Query("to_user_id"), 10, 64); err != nil{
		context.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  "to_user_id格式错误",
		})
		return
	}

	var actionType int64
	if actionType, err = strconv.ParseInt(context.Query("action_type"), 10, 64); err != nil || actionType < 1 || actionType > 2{
		context.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  "action_type格式错误",
		})
		return
	}

	if result, err := service.Follow(userId, toUserId, actionType); result && err == nil{
		context.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "OK",
			})
	}else{
		context.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  "关注/取关操作失败",
		})
	}
}

func FollowListController(context *gin.Context){
	var err error
	var userId int64
	if userId, err = strconv.ParseInt(context.GetString("userId"), 10, 64); err != nil{
		context.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  "user_id格式错误",
		})
		return
	}

	if userList, err := service.FollowList(userId); err == nil{
		context.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "获取关注列表成功",
			UserList: userList,
		})
	}
	context.JSON(http.StatusOK, Response{
		StatusCode: -1,
		StatusMsg:  "获取关注列表失败",
	})
}

func FollowerListController(context *gin.Context){
	var err error
	var userId int64
	if userId, err = strconv.ParseInt(context.GetString("userId"), 10, 64); err != nil{
		context.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  "user_id格式错误",
		})
		return
	}

	if userList, err := service.FollowerList(userId); err == nil{
		context.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "获取粉丝列表成功",
			UserList: userList,
		})
	}
	context.JSON(http.StatusOK, Response{
		StatusCode: -1,
		StatusMsg:  "获取粉丝列表失败",
	})
}
