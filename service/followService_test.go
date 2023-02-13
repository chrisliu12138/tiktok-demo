package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/middleware/rabbitmq"
	"github.com/RaymondCode/simple-demo/middleware/redis"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	var err error
	//先打开redis、rabbitmq、mysql
	err = redis.InitRedis()
	assert.NoError(t, err)

	err = rabbitmq.InitRabbitMQ()
	assert.NoError(t, err)

	err = dao.InitDB()
	assert.NoError(t, err)

	//关注
	result, err := Follow(10, 11, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	//取关
	result, err = Follow(10, 11, 2)
	assert.True(t, result)
	assert.NoError(t, err)

	//确保rabbitmq消息已经发送
	time.Sleep(time.Second)
	
	err = rabbitmq.FollowRmq.Close()
	assert.NoError(t, err)
}

func TestFollowList(t *testing.T) {
	var err error
	//先打开redis、rabbitmq、mysql
	err = redis.InitRedis()
	assert.NoError(t, err)

	err = rabbitmq.InitRabbitMQ()
	assert.NoError(t, err)

	err = dao.InitDB()
	assert.NoError(t, err)

	//关注,确保数据库中至少有一条记录
	result, err := Follow(10, 11, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	user, err := FollowList(10)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	fmt.Println(user)

	//确保rabbitmq消息已经发送
	time.Sleep(time.Second)

	err = rabbitmq.FollowRmq.Close()
	assert.NoError(t, err)
}

func TestFollowerList(t *testing.T) {
	var err error
	//先打开redis、rabbitmq、mysql
	err = redis.InitRedis()
	assert.NoError(t, err)

	err = rabbitmq.InitRabbitMQ()
	assert.NoError(t, err)

	err = dao.InitDB()
	assert.NoError(t, err)

	//关注,确保数据库中至少有一条记录
	result, err := Follow(10, 11, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	user, err := FollowerList(10)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	fmt.Println(user)

	//确保rabbitmq消息已经发送
	time.Sleep(time.Second)

	err = rabbitmq.FollowRmq.Close()
	assert.NoError(t, err)
}