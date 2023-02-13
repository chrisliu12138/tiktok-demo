package service

import (
	"SimpleDouyin/dao"
)

// User controller进行最终封装返回的User用户结构体
//type User struct {
//	Id             int64  `json:"id,omitempty"`
//	Name           string `json:"name,omitempty"`
//	FollowCount    int64  `json:"follow_count"`
//	FollowerCount  int64  `json:"follower_count"`
//	IsFollow       bool   `json:"is_follow"`
//	TotalFavorited int64  `json:"total_favorited,omitempty"`
//	FavoriteCount  int64  `json:"favorite_count,omitempty"`
//}

type UserService interface {
	/*
		本地调用API
	*/
	// GetTableUserList 获得全部TableUser对象
	GetTableUserList() []dao.TableUser

	// GetTableUserByUsername 根据UserName获得TableUser对象
	GetTableUserByUsername(name string) dao.TableUser

	// GetTableUserById 根据user_id获得TableUser对象
	GetTableUserById(id int64) dao.TableUser

	// InsertTableUser 将tableUser对象插入表内
	InsertTableUser(tableUser *dao.TableUser) bool

	/*
		向外暴露API
	*/

	// GetUserById 未登录情况下 根据user_id获得User对象
	GetUserById(id int64) (dao.User, error)

	// GetUserByIdWithCurId 已登录(curId)情况下 根据user_id获得User对象
	GetUserByIdWithCurId(id int64, curId int64) (dao.User, error)
}
