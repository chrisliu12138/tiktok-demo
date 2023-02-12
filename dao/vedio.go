package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Utils"
	//"github.com/RaymondCode/simple-demo/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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
func (v Video) TableName() string {
	return "video"
}

// 上传稿件   userID:是谁发的(根据token去用户信息表查id)  playUrl：存在哪里  title：视频名字是啥
func Add(uerId uint, playUrl string, title string) bool {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      true,            // 彩色打印
		},
	)
	dsn := fmt.Sprintf("root:root123@tcp(1.117.88.168:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local")
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	video := VideoEntity{Title: title, PlayUrl: playUrl, UserID: uerId}
	result := db.Create(&video) // 通过数据的指针来创建
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      true,            // 彩色打印
		},
	)
	dsn := fmt.Sprintf("tiktok:tiktok123@tcp(1.117.88.168:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local")
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	//连接数据库并查询
	// SELECT * FROM `video` left join user on user.id = video.user_id where user_id = usrId;
	result := db.Model(&Video{}).
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
	//连接数据库并查询
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      true,            // 彩色打印
		},
	)
	dsn := fmt.Sprintf("tiktok:tiktok123@tcp(1.117.88.168:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local")
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// SELECT * FROM `video` left join user on user.id = video.user_id where video.id = ID;
	result := db.Model(&Video{}).
		Select("video.id,title,play_url,cover_url,favorite_count,comment_count,is_favorite,video.create_time,user_id,name,follow_count,follower_count,bool").Joins("left join user on user.id = video.user_id").Where(videoArray).Scan(&rows)

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
	//连接数据库并查询
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 2 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      true,            // 彩色打印
		},
	)
	dsn := fmt.Sprintf("tiktok:tiktok123@tcp(1.117.88.168:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local")
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	//连接数据库并查询
	// SELECT * FROM `video` left join user on user.id = video.user_id ORDER BY create_time desc  LIMIT 30;
	result := db.Model(&Video{}).
		Select("video.id,title,play_url,cover_url,favorite_count,comment_count,is_favorite,video.create_time,user_id,name,follow_count,follower_count,bool").Joins("left join user on user.id = video.user_id").Order("video.create_time desc").Limit(30).Scan(&rows)
	if result.Error != nil {
		return nil //查询失败
	}
	return rows
}

// 拿到当前的所有视频id，限制起止数
func GetVedioIdWithLimit(start, limit int64) []int64 {
	var videoIds []int64
	err := Utils.DB.Raw("SELECT ID FROM video OFFSET ? LIMIT ?", start, limit).Scan(&videoIds)
	if err != nil {
		panic(err)
	}
	return videoIds
}

// 返回视频总数
func GetVedioCount() int64 {
	var count int64
	Utils.DB.Raw("SELECT count(1) from video").Scan(&count)
	return count
}

// 更新mysql中的点赞总数
//func UpdateVedioLikeCount(vedioId, likeCount int64) bool {
//	//Video := service.Video{ID: uint(vedioId)}
//	//result := true
//	//err := Utils.DB.Model(&Video).Update("FavoriteCount", likeCount).Error
//	//if err != nil {
//	//	result = false
//	//	panic(err)
//	//}
//	//return result
//}
