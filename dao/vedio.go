package dao

import (
	"SimpleDouyin/middleware/DBUtils"
	"fmt"
	"time"
)

// VideoEntity
type VideoEntity struct {
	ID            uint `gorm:"primarykey"`
	Title         string
	PlayUrl       string `gorm:"default:unknown"`
	CoverUrl      string `gorm:"default:unknown"`
	FavoriteCount uint   `gorm:"default:0"`
	CommentCount  uint   `gorm:"default:0"`
	IsFavorite    uint   `gorm:"default:0"`
	UserID        uint
	CreatedAt     time.Time `gorm:"column:create_time"`
}

// 自定义Entity
type Result struct {
	ID            uint `gorm:"primarykey"`
	Title         string
	PlayUrl       string    `gorm:"default:unknown"`
	CoverUrl      string    `gorm:"default:unknown"`
	FavoriteCount uint      `gorm:"default:0"`
	CommentCount  uint      `gorm:"default:0"`
	IsFavorite    uint      `gorm:"default:0"`
	UserID        uint      `gorm:"column:user_id"`
	Name          string    `gorm:"column:name"`
	FollowCount   uint      `gorm:"column:follow_count"`
	FollowerCount uint      `gorm:"column:follower_count"`
	IsFollow      uint      `gorm:"column:bool"`
	CreatedAt     time.Time `gorm:"column:create_time"`
}

// 定义表名
func (v VideoEntity) TableName() string {
	return "video"
}

// 上传稿件   userID:是谁发的(根据token去用户信息表查id)  playUrl：存在哪里  title：视频名字是啥
func Add(uerId uint, playUrl string, title string) bool {
	DBUtils.InitMysqlTemplete()
	video := VideoEntity{Title: title, PlayUrl: playUrl, UserID: uerId}
	result := DBUtils.DB.Create(&video) // 通过数据的指针来创建
	if result.Error != nil {
		return false //上传失败
	}

	fmt.Println(video.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	return true //上传成功
	//fmt.Println(video.ID)            // 返回插入数据的主键
	//fmt.Println(result.Error)        // 返回 error
	//fmt.Println(result.RowsAffected) // 返回插入记录的条数
}

// 根据userID查询稿件
func Query(userid uint) []Result {
	//查询某个用户的所有视频
	//rows := make([]*Result, 0)
	var rows []Result
	//连接数据库并查询
	// SELECT * FROM `video` left join user on user.id = video.user_id where user_id = usrId;
	result := DBUtils.DB.Model(&VideoEntity{}).
		Select("video.id,title,play_url,cover_url,favorite_count,comment_count,is_favorite,video.create_time,user.id,name,follow_count,follower_count,bool").Joins("left join user on user.id = video.user_id").Where("user.id = ?", userid).Scan(&rows)
	if result.Error != nil {
		return nil //查询失败
	}
	return rows

}

// 根据videoArray查询稿件
func QueryListByVedionl(videoArray []int64) []Result {
	//rows := make([]*Result, 0)
	var rows []Result
	DBUtils.InitMysqlTemplete()
	result := DBUtils.DB.Model(&VideoEntity{}).
		Select("video.id,video.title,video.play_url,video.cover_url,video.favorite_count,video.comment_count,video.is_favorite,video.create_time,video.user_id,user.name,user.follow_count,user.follower_count,user.bool").
		Joins("left join user on user.id = video.user_id").Where("video.id in ?", videoArray).Scan(&rows)

	if result.Error != nil {
		return nil //查询失败
	}
	return rows

}

// 查询最新的30个稿件
func QueryAll() []Result {
	//查询最新的30条数据
	//rows := make([]*Result, 0)
	var rows []Result
	DBUtils.InitMysqlTemplete()
	//连接数据库并查询
	// SELECT * FROM `video` left join user on user.id = video.user_id ORDER BY create_time desc  LIMIT 30;
	result := DBUtils.DB.Model(&VideoEntity{}).
		Select("video.id,title,play_url,cover_url,favorite_count,comment_count,is_favorite,video.create_time,user_id,name,follow_count,follower_count,bool").Joins("left join user on user.id = video.user_id").Order("video.create_time desc").Limit(30).Scan(&rows)
	if result.Error != nil {
		return nil //查询失败
	}
	return rows
}
func Sqltest() {
	var rows []Result

	result := DBUtils.DB.Model(&Video{}).
		Select("video.id,title,video.create_time,user.id").Joins("left join user on user.id = video.user_id").Order("video.create_time desc").Limit(30).Scan(&rows)
	fmt.Println(result)

}

// 拿到当前的所有视频id，限制起止数
func GetVedioIdWithLimit(start, limit int64) []int64 {
	var videoIds []int64
	err := DBUtils.DB.Raw("SELECT ID FROM video LIMIT ? OFFSET ?", limit, start).Scan(&videoIds).Error
	if err != nil {
		panic(err)
	}
	return videoIds
}

// 返回视频总数
func GetVedioCount() int64 {
	var count int64
	err := DBUtils.DB.Raw("SELECT count(1) from video").Find(&count).Error
	if err != nil {
		panic(err)
	}
	return count
}

// 更新mysql中的点赞总数
func UpdateVedioLikeCount(vedioId, likeCount int64) bool {
	Video := VideoEntity{ID: uint((vedioId))}
	result := true
	err := DBUtils.DB.Table("video").Model(&Video).Update("FavoriteCount", likeCount).Error
	if err != nil {
		result = false
		panic(err)
	}
	return result
}
