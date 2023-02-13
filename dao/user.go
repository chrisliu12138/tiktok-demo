package dao

import (
	"SimpleDouyin/middleware/DBUtils"
	"log"
)

/**
用户实体类及相关操作数据库方法
*/

// TableUser 对应数据库User表结构的结构体
type TableUser struct {
	Id       int    `gorm:"primarykey"`
	Name     string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

// TableName 修改表名映射
func (tableUser TableUser) TableName() string {
	return "user"
}

// GetTableUserList 获取全部TableUser对象
func GetTableUserList() ([]TableUser, error) {
	var tableUsers []TableUser
	err := DBUtils.DB.Find(&tableUsers).Error
	if err != nil {
		log.Println(err.Error())
		return tableUsers, err
	}
	return tableUsers, nil
}

// GetTableUserByUserName 根据username获得TableUser对象
func GetTableUserByUserName(name string) (TableUser, error) {
	tableUser := TableUser{}
	err := DBUtils.DB.Where("name = ?", name).First(&tableUser).Error
	if err != nil {
		log.Println(err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// GetTableUserById 根据user_id获得TableUser对象
func GetTableUserById(id int64) (TableUser, error) {
	tableUser := TableUser{}
	err := DBUtils.DB.Where("id = ?", id).First(&tableUser).Error
	if err != nil {
		log.Println(err.Error())
		return tableUser, err
	}
	return tableUser, err
}

// InsertTableUser 将tableUser插入表中
func InsertTableUser(tableUser *TableUser) bool {
	err := DBUtils.DB.Create(&tableUser).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
