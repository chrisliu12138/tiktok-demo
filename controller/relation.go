package controller

import (
	"SimpleDouyin/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	dao.Response
	UserList []dao.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, dao.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: dao.Response{
			StatusCode: 0,
		},
		UserList: []dao.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: dao.Response{
			StatusCode: 0,
		},
		UserList: []dao.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: dao.Response{
			StatusCode: 0,
		},
		UserList: []dao.User{DemoUser},
	})
}
