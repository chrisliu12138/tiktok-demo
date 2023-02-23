package service

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/middleware/DBUtils"
	"errors"
	"gorm.io/gorm"
	"log"
)

// Comment 新增评论
func Comment(videoId int64, userId int64, commentText string) (comment dao.Comment, err error) {

	// 开启事务，以便获取行锁
	tx := DBUtils.DB.Begin()
	// 判断video是否存在
	var videoCommentCount int64

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return comment, errors.New("评论的视频不存在")
	}
	// 增加视频评论数并更新
	videoCommentCount++
	if err = tx.Model(&dao.Video{}).Where("targetId = ?", videoId).
		Update("comment_count", videoCommentCount).Error; err != nil {
		tx.Rollback()
		return comment, err
	}
	tx.Commit()
	c := dao.Comment{Content: commentText}
	// 新建comment对象并保存在数据库中
	dao.InsertComment(c)
	return comment, nil
}

// DeleteComment 删除评论
func DeleteComment(commentId int64) (err error) {

	// 判断删除的评论是否存在
	err = dao.DeleteComment(commentId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("删除的评论不存在")
	}
	return err
}

// GetCommentList 获取评论列表
func GetCommentList(videoId int64) ([]dao.Comment, error) {
	CommentList, err := dao.GetCommentList(videoId)
	if err != nil {
		log.Println("Err:", err.Error())
		return CommentList, nil
	}
	return CommentList, nil
}
