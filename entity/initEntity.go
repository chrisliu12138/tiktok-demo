package entity

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

var RDB *redis.Client

func Init() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold: time.Second,  //慢SQL阈值
			LogLevel:      logger.Error, //Log Level
			Colorful:      true,         //彩色打印
		},
	)
	var err error
	dsn := "root:root123@tcp(1.117.88.168:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic("err:", err.Error())
	}
}

func GetRedisTemplete() *redis.Client {
	if RDB != nil {
		return RDB
	}
	opt, err := redis.ParseURL("redis://1.117.88.168/:6379/0")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}
