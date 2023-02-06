package dao

import (
	"github.com/RaymondCode/simple-demo/Utils"
	"log"
)

// TableUser 对应数据库User表结构的结构体
type TableUser struct {
	Id       int64
	Name     string
	Password string
}

// TableName 修改表名映射
func (tableUser TableUser) TableName() string {
	return "users"
}

// GetTableUserList 获取全部TableUser对象
func GetTableUserList() ([]TableUser, error) {
	var tableUsers []TableUser
	err := Utils.DB.Find(&tableUsers).Error
	if err != nil {
		log.Println(err.Error())
		return tableUsers, err
	}
	return tableUsers, nil
}

// GetTableUserByUserName 根据username获得TableUser对象
func GetTableUserByUserName(name string) (TableUser, error) {
	tableUser := TableUser{}
	err := Utils.DB.Where("name = ?", name).First(&tableUser).Error
	if err != nil {
		log.Println(err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// GetTableUserById 根据user_id获得TableUser对象
func GetTableUserById(id int64) (TableUser, error) {
	tableUser := TableUser{}
	err := Utils.DB.Where("id = ?", id).First(&tableUser).Error
	if err != nil {
		log.Println(err.Error())
		return tableUser, err
	}
	return tableUser, err
}

// InsertTableUser 将tableUser插入表中
func InsertTableUser(tableUser *TableUser) bool {
	err := Utils.DB.Create(&tableUser).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
