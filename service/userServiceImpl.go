package service

import (
	"SimpleDouyin/config"
	"SimpleDouyin/dao"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	/*
		FollowService
		LikeService
	*/
}

// GetTableUserList 获得全部TableUser对象
func (userServiceImpl *UserServiceImpl) GetTableUserList() []dao.TableUser {
	tableUsers, err := dao.GetTableUserList()
	if err != nil {
		log.Println("Err:", err.Error())
		return tableUsers
	}
	return tableUsers
}

// GetTableUserByUsername 根据Username获得TableUser对象
func (userServiceImpl *UserServiceImpl) GetTableUserByUsername(name string) dao.TableUser {
	tableUser, err := dao.GetTableUserByUserName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		return tableUser
	}
	return tableUser
}

// GetTableUserById 根据user_id获得TableUser对象
func (userServiceImpl *UserServiceImpl) GetTableUserById(id int64) dao.TableUser {
	tableUser, err := dao.GetTableUserById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		return tableUser
	}
	return tableUser
}

// InsertTableUser 将tableUser对象插入表内
func (userServiceImpl *UserServiceImpl) InsertTableUser(tableUser *dao.TableUser) bool {
	flag := dao.InsertTableUser(tableUser)
	if flag == false {
		log.Println("新增用户插入失败")
		return false
	}
	return true
}

// GetUserById 未登录情况下 根据user_id获得User对象
func (userServiceImpl *UserServiceImpl) GetUserById(id int64) (dao.User, error) {
	user := dao.User{
		Id:            0,
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	tableUser, err := dao.GetTableUserById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user, err
	}
	log.Println("User Query Success")
	followCnt, err := GetFollowCnt(id)
	if err != nil {
		log.Println("Err:", err.Error())
	}
	folloerCnt, err := GetFollowerCnt(id)
	if err != nil {
		log.Println("Err:", err.Error())
	}
	totalFavorited := GetUserLikeCount(string(id))
	favoriteCnt := len(GetVedioLikeList(string(id)))
	user = dao.User{
		Id:             id,
		Name:           tableUser.Name,
		FollowCount:    followCnt,
		FollowerCount:  folloerCnt,
		IsFollow:       false,
		TotalFavorited: totalFavorited,
		FavoriteCount:  int64(favoriteCnt),
	}
	return user, nil
}

// GetUserByIdWithCurId 已登录(curId)情况下 根据user_id获得User对象
func (userServiceImpl *UserServiceImpl) GetUserByIdWithCurId(id int64, curId int64) (dao.User, error) {
	user := dao.User{
		Id:            0,
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false, //数据库中存的 0：表示false 1 表示成功
	}
	tableUser, err := dao.GetTableUserById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user, err
	}
	log.Println("User Query Success")
	followCnt, err := GetFollowCnt(id)
	if err != nil {
		log.Println("Err:", err.Error())
	}
	folloerCnt, err := GetFollowerCnt(id)
	if err != nil {
		log.Println("Err:", err.Error())
	}
	totalFavorited, err := TotalFavourite(id)
	if err != nil {
		log.Println("Err:", err.Error())
	}
	favoriteCnt := len(GetVedioLikeList(string(id)))
	user = dao.User{
		Id:             id,
		Name:           tableUser.Name,
		FollowCount:    followCnt,
		FollowerCount:  folloerCnt,
		IsFollow:       false,
		TotalFavorited: totalFavorited,
		FavoriteCount:  int64(favoriteCnt),
	}
	return user, nil
}

// GenerateToken 根据userName生成一个token
func GenerateToken(userName string) string {
	user := UserService.GetTableUserByUsername(new(UserServiceImpl), userName)
	token := NewToken(user)
	log.Printf("GenerateToken: %v\n\n", token)
	return token
}

// NewToken 根据信息创建Token
func NewToken(user dao.TableUser) string {
	expireTime := time.Now().Unix() + int64(config.ONE_DAY_HOUR)
	log.Printf("ExpireTime is %v\n\n", expireTime)
	claims := jwt.StandardClaims{
		Audience:  user.Name,
		ExpiresAt: expireTime,
		Id:        strconv.FormatInt(int64(user.Id), 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.SECRET)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err == nil {
		token = "Bobochang" + token
		log.Println("Generate token success!")
		return token
	} else {
		log.Println("Generate token fail")
		return "fail"
	}
}

// EnCoder 密码加密
func EnCoder(password string) string {
	hash := hmac.New(sha256.New, []byte(password))
	encryptPwd := hex.EncodeToString(hash.Sum(nil))
	log.Println("EncryptPassword is ", encryptPwd)
	return encryptPwd
}
