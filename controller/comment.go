package controller

import (
	"net/http"

	"SimpleDouyin/dao"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	dao.Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	dao.Response
	Comment dao.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//token := c.Query("token")
	//actionType := c.Query("action_type")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	if actionType == "1" {
	//		text := c.Query("comment_text")
	//		c.JSON(http.StatusOK, CommentActionResponse{Response: dao.Response{StatusCode: 0},
	//			Comment: dao.Comment{
	//				Id:         1,
	//				User:       user,
	//				Content:    text,
	//				CreateDate: "05-01",
	//			}})
	//		return
	//	}
	//	c.JSON(http.StatusOK, dao.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, dao.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    dao.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}

//下面的两个方法写在service层
func addComment(vedioId int64, comment string) bool {

	//判断是否存在
	/*

		count := db.raw.(select count(1) from comment )

		if count > 0{
		//错误，已存在

		}else{

			1.新建comment对象，使用前端的参数赋值

			2.向mysql中添加数据

		}

	*/
	return true

}

func deleteComment() {
	//逻辑如上，先判断在删除
}
