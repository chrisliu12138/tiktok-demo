package main

import (
<<<<<<< HEAD
	"github.com/RaymondCode/simple-demo/service"
=======
	"github.com/RaymondCode/simple-demo/entity"
>>>>>>> master
	"github.com/gin-gonic/gin"
)

func main() {
<<<<<<< HEAD
	go service.RunMessageServer()

	r := gin.Default()

=======
	// go service.RunMessageServer()

	InitDeps()

	//gin
	r := gin.Default()
>>>>>>> master
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
<<<<<<< HEAD
=======

// InitDeps 加载项目依赖
func InitDeps() {
	// 初始化数据库连接
	entity.Init()
}
>>>>>>> master
