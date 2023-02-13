package serviceTest

import (
	"SimpleDouyin/service"
	"fmt"
	"testing"
)

func TestQueryListByVedionl(t *testing.T) {
	impl := service.VideoServiceImpl{}
	videoList := impl.QueryListByVedioIdList([]int64{1, 2, 3, 4})
	if videoList == nil {
		t.Error("测试失败")
	} else {
		fmt.Print("videolist is that:", videoList)
		t.Logf("测试成功")
	}

}

func TestQuery(t *testing.T) {
	impl := service.VideoServiceImpl{}
	videoList := impl.Query(8)
	if videoList == nil {
		t.Error("测试失败")
	} else {
		fmt.Print("videolist is that:", videoList)
		t.Logf("测试成功")
	}

}

func TestQueryAll(t *testing.T) {
	impl := service.VideoServiceImpl{}
	videoList := impl.QueryAll()
	if videoList == nil {
		t.Error("测试失败")
	} else {
		fmt.Print("videolist is that:", videoList)
		t.Logf("测试成功")
	}

}
