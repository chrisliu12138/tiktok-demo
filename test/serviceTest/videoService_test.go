package serviceTest

import (
	"SimpleDouyin/dao"
	"testing"
)

func TestQueryListByVedionl(t *testing.T) {

	result := dao.QueryListByVedionl([]int64{1, 2, 3, 4})
	if result == nil {
		t.Error("测试失败")
	}
	t.Logf("测试成功")

}

func TestQuery(t *testing.T) {
	result := dao.Query(8)
	if result == nil {
		t.Error("测试失败")
	}
	t.Logf("测试成功")

}
func TestQueryAll(t *testing.T) {
	result := dao.QueryAll()
	if result == nil {
		t.Error("测试失败")
	}
	t.Logf("测试成功")

}
