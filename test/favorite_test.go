package test

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Utils"
	"github.com/RaymondCode/simple-demo/service"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var name = []string{"1", "2", "3", "4"}

/*
Tips：
单元测试需要在方法名前加大写开头的Test
需要在方法中传入测试的固定参数t *testing.T
可以使用log

基准测试相当于简易版的性能测试，会显示接口的执行速度，执行次数的信息

建议手动点一下，看看控制台
*/
func TestGetVedioLikeList(t *testing.T) {
	Utils.InitRedisTemplete()
	if Utils.RDB == nil {
		fmt.Println("初始化失败")
	}
	list := service.GetVedioLikeCount("2")
	fmt.Print("list is that:", list)
	service.GetVedioLikeList("2")

}

func TestAdddata(t *testing.T) {
	Utils.InitRedisTemplete()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Int63n(30)
		service.AdduserId(strconv.Itoa(2), strconv.Itoa(int(num)))
	}
}

func BenchmarkDislikeVedio(b *testing.B) {
	Utils.InitRedisTemplete()
	result := service.DislikeVedio("11", "15")
	fmt.Println(result)
}

func TestTimeClock(t *testing.T) {
	Utils.InitRedisTemplete()
	Utils.TimeMission()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Int63n(30)
		fmt.Println(num)
	}

}
