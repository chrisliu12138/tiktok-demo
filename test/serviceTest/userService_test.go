package serviceTest

import (
	"SimpleDouyin/middleware/DBUtils"
	"SimpleDouyin/service"
	"fmt"
	"testing"
)

func TestGetTableUserList(t *testing.T) {
	impl := service.UserServiceImpl{}
	list := impl.GetTableUserList()
	fmt.Printf("%v", list)
}

func TestGetTableUserByUsername(t *testing.T) {
	DBUtils.Init()
	impl := service.UserServiceImpl{}
	list := impl.GetTableUserByUsername("aaa")
	fmt.Printf("%v", list)
}

func TestGetTableUserById(t *testing.T) {
	impl := service.UserServiceImpl{}
	list := impl.GetTableUserById(int64(4))
	fmt.Printf("%v", list)
}

//func TestInsertTableUser(t *testing.T) {
//	impl := service.UserServiceImpl{}
//	user := &impl.TableUser{
//		Id:       20000,
//		Name:     "qaq",
//		Password: "111111",
//	}
//	list := impl.InsertTableUser(user)
//	fmt.Printf("%v", list)
//}

func TestGetUserById(t *testing.T) {
	impl := service.UserServiceImpl{
		//FollowService: &FollowServiceImp{},
		//LikeService:   &LikeServiceImpl{},
	}
	list, _ := impl.GetUserById(int64(4))
	fmt.Printf("%v", list)
}

func TestGetUserByIdWithCurId(t *testing.T) {
	impl := service.UserServiceImpl{
		//FollowService: &FollowServiceImp{},
	}
	list, _ := impl.GetUserByIdWithCurId(int64(482), int64(130))
	fmt.Printf("%v", list)
}
