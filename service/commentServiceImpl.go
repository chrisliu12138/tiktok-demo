package service

import (
	"SimpleDouyin/dao"
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Comment 新增评论
func Comment(videoId int64, userId int64, commentText string) (comment dao.Comment, err error) {
	// 判断进行评论的用户是否存在
	user, err := UserInfo(userId)
	if err != nil {
		return comment, errors.New("执行评论操作的用户不存在")
	}
	// 开启事务，以便获取行锁
	tx := dao.DB.Begin()
	// 判断video是否存在
	err := dao.InsertComment(comment)
	var videoCommentCount int64

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return comment, errors.New("评论的视频不存在")
	}
	// 增加视频评论数并更新
	videoCommentCount++
	if err = tx.Model(&dao.Video{}).Where("id = ?", targetId).
		Update("comment_count", videoCommentCount).Error; err != nil {
		tx.Rollback()
		return comment, err
	}
	tx.Commit()

	// 新建comment对象并保存在数据库中
	comment.UserId = user.Id
	comment.TargerId = targetId
	comment.Content = commentText
	comment.CreateDate = time.Now().Format("01-02")
	comment.User = user
	if err = dao.DB.Model(&dao.Comment{}).Create(&comment).Error; err != nil {
		return comment, err
	}
	tx.Commit()
	return comment, nil
}

// DeleteComment 删除评论
func DeleteComment(commentId int64) (err error) {
	comment := dao.Comment{}

	// 判断删除的评论是否存在
	err = dao.DeleteComment(comment)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("删除的评论不存在")
	}

	// 判断删除评论所在的视频是否存在
	// 开启事务，以便获取行锁
	tx := dao.DB.Begin()
	// 判断video是否存在
	var videoCommentCount int64
	err = tx.Model(&dao.Video{}).Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", comment.targetId).Select("comment_count").Find(&videoCommentCount).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return errors.New("当前评论所在的视频不存在")
	}
	// 如果当前视频评论数不为正数，则删除评论操作是异常的，会导致视频评论数变为负数
	if videoCommentCount <= 0 {
		tx.Rollback()
		return errors.New("当前评论所在的视频评论数异常")
	}
	// 减少视频评论数并更新
	videoCommentCount--
	if err = tx.Model(&dao.Video{}).Where("id = ?", comment.targetId).
		Update("comment_count", videoCommentCount).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	// 删除在评论表中的对应评论行数据
	if err = dao.DB.Model(&dao.Comment{}).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

// GetCommentList 获取评论列表
func (commentServiceImpl *CommentServiceImpl) GetCommentList(videoId uint) ([]model.Comment, error) {
	()[]dao.TableUser {
		CommentList, err := dao.GetCommentList()
	if err != nil {
		log.Println("Err:", err.Error())
		return CommentList
	}
	return CommentList
}