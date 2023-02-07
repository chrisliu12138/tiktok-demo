package main

import (
	"github.com/RaymondCode/simple-demo/Utils"
	"github.com/gin-gonic/gin"
)

func main() {

	//go service.RunMessageServer()

	r := gin.Default()

	// go service.RunMessageServer()

	InitDeps()

	//gin
	r = gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// InitDeps 加载项目依赖
func InitDeps() {
	// 初始化数据库连接
	Utils.Init()
	Utils.TimeMission()
}
