package test

import (
	"SimpleDouyin/middleware/ffmpeg"
	"SimpleDouyin/middleware/ftp"
	"testing"
)

// 生成一张截图
func TestFfmpeg(t *testing.T) {
	ffmpeg.InitSSH()
	//输入：服务器目录下的bear
	//输入:设置生成的截图名称
	ffmpeg.Ffmpeg("bear", "output")
}

// 上传视频到服务器
func TestFtp(t *testing.T) {
	//相对路径：指的是调用Ftp函数的位置(不是FTP函数所在的位置)
	ftp.Ftp("../public/bear1.mp4")
}
