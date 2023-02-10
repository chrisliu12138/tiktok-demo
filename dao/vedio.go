package dao

import (
	"github.com/RaymondCode/simple-demo/Utils"
	"github.com/RaymondCode/simple-demo/service/impl"
)

//拿到当前的所有视频id，限制起止数
func GetVedioIdWithLimit(start, limit int64) []int64 {
	var videoIds []int64
	err := Utils.DB.Raw("SELECT ID FROM video OFFSET ? LIMIT ?", start, limit).Scan(&videoIds)
	if err != nil {
		panic(err)
	}
	return videoIds
}

//返回视频总数
func GetVedioCount() int64 {
	var count int64
	Utils.DB.Raw("SELECT count(1) from video").Scan(&count)
	return count
}

//更新mysql中的点赞总数
func UpdateVedioLikeCount(vedioId, likeCount int64) bool {
	Video := impl.Video{ID: uint(vedioId)}
	result := true
	err := Utils.DB.Model(&Video).Update("FavoriteCount", likeCount).Error
	if err != nil {
		result = false
		panic(err)
	}
	return result
}
