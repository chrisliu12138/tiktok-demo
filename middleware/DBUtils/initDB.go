package DBUtils

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var RDB *redis.Client

func InitMysqlTemplete() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold: time.Second,  //慢SQL阈值
			LogLevel:      logger.Error, //Log Level
			Colorful:      true,         //彩色打印
		},
	)
	var err error
	dsn := "tiktok:tiktok123@tcp(1.117.88.168:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic("err:", err.Error())
	}
}

func InitRedisTemplete() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "1.117.88.168:6379",
		Password: "123321", // 密码
		DB:       0,        // 数据库，从0开始
		PoolSize: 30,       // 连接池大小
	})
}

func Init() {
	InitRedisTemplete()
	InitMysqlTemplete()
}
