package rabbitmq

import (
	"fmt"
	"log"
	"strings"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/middleware/redis"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
	queueName string
}

var FollowRmq RabbitMQ
var UnFollowRmq RabbitMQ

func (r *RabbitMQ) Connect(name string) error{
	var err error
	r.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("[amqp] connect error: %s\n", err)
		return err
	}
	r.channel, err = r.conn.Channel()
	if err != nil {
		log.Printf("[amqp] get channel error: %s\n", err)
		return err
	}
	r.done = make(chan error)
	r.queueName = name
	return nil
}

func (r *RabbitMQ) Publish(body string){
	_, err := r.channel.QueueDeclare(
		r.queueName, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	err = r.channel.Publish(
		"",     // exchange
		r.queueName, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			//消息持久化
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func (r *RabbitMQ) Consume(){
	_, err := r.channel.QueueDeclare(
    r.queueName, // name
    true,   // durable
    false,   // delete when usused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")

  msgs, err := r.channel.Consume(
    r.queueName, // queue
    "",     // consumer
    false,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
  )
  failOnError(err, "Failed to register a consumer")
  
  forever := make(chan bool)

	switch r.queueName {
	case "follow":
		r.follow(msgs)
	case "unfollow":
		r.unfollow(msgs)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}


func (r *RabbitMQ) follow(msgs <-chan amqp.Delivery){
	go func() {
    for d := range msgs {
      log.Printf("Received a message: %s", d.Body)
			params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
			userIdStr := params[0]
			toUserIdStr := params[1]

			//redis更新
			//1、user的关注列表+1
			//2、target_user的粉丝列表+1

			//redis关注列表更新
			if num, err := redis.RdbFollow.SCard(redis.Ctx, userIdStr).Result(); num != 0 && err == nil{
				redis.RdbFollow.SAdd(redis.Ctx, userIdStr, toUserIdStr)
				redis.RdbFollow.Expire(redis.Ctx, userIdStr, config.ExpireTime)
			}else if err != nil{
				log.Println("关注列表更新错误: ", err)
				return
			}else{
				
			}

			//redis粉丝列表更新
			if num, err := redis.RdbFollower.SCard(redis.Ctx, toUserIdStr).Result(); num != 0 && err == nil{
				redis.RdbFollower.SAdd(redis.Ctx, toUserIdStr, userIdStr)
				redis.RdbFollower.Expire(redis.Ctx, toUserIdStr, config.ExpireTime)
			}else if err != nil{
				log.Println("粉丝列表更新错误: ", err)
				return
			}
			//确保消息确实执行完成
			d.Ack(false)
		}
  }()
}

func (r *RabbitMQ) unfollow(msgs <-chan amqp.Delivery){
	go func() {
    for d := range msgs {
      log.Printf("Received a message: %s", d.Body)
			params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
			userIdStr := params[0]
			toUserIdStr := params[1]

			//redis关注列表更新
			if num, err := redis.RdbFollow.SCard(redis.Ctx, userIdStr).Result(); num != 0 && err == nil{
				redis.RdbFollower.SRem(redis.Ctx, userIdStr, toUserIdStr)
				redis.RdbFollower.Expire(redis.Ctx, userIdStr, config.ExpireTime)
			}else if err != nil{
				log.Println("关注列表更新错误: ", err)
				return
			}

			//redis粉丝列表更新
			if num, err := redis.RdbFollower.SCard(redis.Ctx, userIdStr).Result(); num != 0 && err == nil{
				redis.RdbFollower.SRem(redis.Ctx, toUserIdStr, userIdStr)
				redis.RdbFollower.Expire(redis.Ctx, toUserIdStr, config.ExpireTime)
			}else if err != nil{
				log.Println("粉丝列表更新错误: ", err)
				return
			}

			//确保消息确实执行完成
			d.Ack(false)
		}
  }()
}

func (r *RabbitMQ) Close() (err error) {
	err = r.conn.Close()
	if err != nil {
		log.Printf("[amqp] close error: %s\n", err)
		return err
	}
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitRabbitMQ() error{
	if err := FollowRmq.Connect("follow"); err != nil{
		return err
	}
	go FollowRmq.Consume()

	if err := UnFollowRmq.Connect("unfollow"); err != nil{
		return err
	}
	go UnFollowRmq.Consume()
	return nil
}