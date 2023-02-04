package controller

import "github.com/go-redis/redis/v8"

var rdb *redis.Client

func getRedisTemplete() *redis.Client {
	opt, err := redis.ParseURL("redis://@http://1.117.88.168/:6379/<db>")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}
