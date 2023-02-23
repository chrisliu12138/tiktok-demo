package service

import "SimpleDouyin/dao"

type CommentService interface {
	/*
		本地调用API
	*/

	// 1.根据targetId获取视频评论数量的接口
	CountFromtargetId(id int64) (int64, error)

	// 2、发表评论，传进来评论的基本信息，返回保存是否成功的状态描述
	Comment(videoId int64, userId int64, commentText string) (comment dao.Comment, err error)

	// 3、删除评论，传入评论id即可，返回错误状态信息
	DeleteComment(commentId int64) (err error)

	// 4、查看评论列表-返回评论list-在controller层再封装外层的状态信息
	GetCommentList(videoId uint) ([]int64, error)
}
