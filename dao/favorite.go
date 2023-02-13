package dao

import (
	"SimpleDouyin/config"
	"SimpleDouyin/middleware/DBUtils"
	"context"

	"github.com/go-redis/redis/v8"
	"strconv"
)

var Vedio_like = config.Vedio_like
var User_like = config.User_like
var limit_ip = config.LIMIT_IP
var time_out = config.LIMIT_PERIOD

var ctx = context.Background()

func SAdd(key string, value string) int64 {
	result, err := DBUtils.RDB.SAdd(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}
	return result
}
func Sremove(key, value string) int64 {
	result, err := DBUtils.RDB.SRem(ctx, Vedio_like+key, value).Result()
	if err != nil {
		panic(err)
	}
	return result
}
func SIsMember(key, value string) bool {
	result, err := DBUtils.RDB.SIsMember(ctx, Vedio_like+key, value).Result()
	if err != nil {
		panic(err)
	}
	return result
}
func Sget(key, value string) {

}
func SMembers(key string) []string {
	result, err := DBUtils.RDB.SMembers(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func LimitIP(ip, vedioId string) bool {
	ctx := context.Background()
	key := limit_ip + ip + vedioId
	result, err := DBUtils.RDB.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			//如果是第一次访问，则set并设置过期时间
			DBUtils.RDB.Set(ctx, key, 0, time_out).Result()
			return true
		} else {
			panic(err)
		}
	}
	count, _ := strconv.ParseInt(result, 10, 64)
	if count <= 10 {
		DBUtils.RDB.Incr(ctx, key)
	} else {
		//过期
		return false
	}
	return true
}

func SelectCount(key string) int64 {
	var count = int64(0)
	result, err := DBUtils.RDB.SCard(ctx, Vedio_like+key).Result()
	if err != nil {
		panic(err)
	}
	count = result
	return count
}
