package serviceTest

import (
	"SimpleDouyin/middleware/DBUtils"
	"SimpleDouyin/service"
	"fmt"
	"testing"
	"time"

	"SimpleDouyin/middleware/rabbitmq"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	var err error
	//先打开redis、rabbitmq、mysql
	err = DBUtils.InitRedis()
	assert.NoError(t, err)

	err = rabbitmq.InitRabbitMQ()
	assert.NoError(t, err)

	err = DBUtils.InitDB()
	assert.NoError(t, err)

	//关注
	result, err := service.Follow(10, 11, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	//取关
	result, err = service.Follow(10, 11, 2)
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
	err = DBUtils.InitRedis()
	assert.NoError(t, err)

	err = rabbitmq.InitRabbitMQ()
	assert.NoError(t, err)

	err = DBUtils.InitDB()
	assert.NoError(t, err)

	//关注,确保数据库中至少有一条记录
	result, err := service.Follow(10, 11, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	user, err := service.FollowList(10)
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
	err = DBUtils.InitRedis()
	assert.NoError(t, err)

	err = rabbitmq.InitRabbitMQ()
	assert.NoError(t, err)

	err = DBUtils.InitDB()
	assert.NoError(t, err)

	//关注,确保数据库中至少有一条记录
	result, err := service.Follow(10, 11, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	user, err := service.FollowerList(10)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	fmt.Println(user)

	//确保rabbitmq消息已经发送
	time.Sleep(time.Second)

	err = rabbitmq.FollowRmq.Close()
	assert.NoError(t, err)
}
