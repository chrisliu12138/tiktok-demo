package config

import "time"

// SECRET 密钥
var SECRET = "tiktok"

// ONE_MINUTE 关于以秒为单位的一天内的各种时间
var ONE_MINUTE = 60
var ONE_HOUR = 60 * 60
var ONE_DAY_HOUR = 60 * 60 * 24
var ONE_MONUTH = 60 * 60 * 24 * 30
var ONE_YEAR = 365 * 60 * 60 * 24

//redis 点赞前缀
const Vedio_like = "vedio_like_"
const User_like = "user_like_"

//定时任务的mysql起止常量
const MYSQL_LIMIT = 999

//定时任务的规定时间
const UPDATE_PERIOD = 2 * time.Hour
