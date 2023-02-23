package ftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
)

// 相对路径
func Ftp(pathfile string) {
	// 连接到 FTP 服务器
	conn, err := ftp.Dial("192.168.111.132:21")
	// ftp.example.com:21是一个模拟的FTP服务器地址，实际使用中需要替换成真正的FTP服务器的地址。
	if err != nil {
		fmt.Println("Error connecting to FTP server:", err)
		return
	}
	defer conn.Quit()

	// 登录到 FTP 服务器
	// 需要提前设置好username and password
	if err := conn.Login("limingjie", "neil66188."); err != nil {
		fmt.Println("Error logging in to FTP server:", err)
		return
	}

	// 打开本地文件(相对路径)
	file, err := os.Open(pathfile)
	if err != nil {
		fmt.Println("Error opening local file:", err)
		return
	}
	defer file.Close()
	fmt.Println(file.Name())
	// 上传文件到 FTP 服务器（文件名）
	if err := conn.Stor("bear.mp4", file); err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
	fmt.Println("File uploaded successfully!!!")
}
