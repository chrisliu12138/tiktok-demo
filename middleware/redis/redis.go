package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//关注列表
var RdbFollow *redis.Client
//粉丝列表
var RdbFollower *redis.Client

// var Ctx context.Context
var Ctx = context.Background()

func InitRedis() error{
	RdbFollow = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "oooooo",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	RdbFollower = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.7:6379", 
		Password: "oooooo",               
		DB:       1,                
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	if _, err := RdbFollow.Ping(context.Background()).Result(); err != nil{
		return err
	}
	if _, err := RdbFollower.Ping(context.Background()).Result(); err != nil{
		return err
	}
	return nil
}