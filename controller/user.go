package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// Register POST /douyin/user/register/ 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userServiceImpl := service.UserServiceImpl{}
	user := userServiceImpl.GetTableUserByUserName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			},
		})
	} else {
		insertUser := entity.TableUser{
			Name:     username,
			Password: service.EnCoder(password),
		}
		if userServiceImpl.InsertTableUser(&insertUser) != true {
			log.Println("Insert User Fail")
		}
		user := userServiceImpl.GetTableUserByUserName(username)
		token := service.GenerateToken(username)
		log.Println("当前用户注册的ID是 ", user.Id)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

// Login POST /douyin/user/login/ 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := service.EnCoder(password)
	log.Println("EncoderPassword is ", encoderPassword)

	userServiceImpl := service.UserServiceImpl{}
	user := userServiceImpl.GetTableUserByUserName(username)
	if encoderPassword == user.Password {
		token := service.GenerateToken(username)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Username or Password Error",
			},
		})
	}
}

// UserInfo GET /douyin/user/ 用户信息
func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)

	userServiceImpl := service.UserServiceImpl{}
	user, err := userServiceImpl.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User Doesn't Exist",
			},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	}
}
