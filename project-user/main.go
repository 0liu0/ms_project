package main

import (
	"github.com/gin-gonic/gin"
	srv "liuche.com/project-common"
	"liuche.com/project-user/router"
)

func main() {
	r := gin.Default()
	// 初始化路由
	router.InitRouter(r)
	srv.Run(r, "userWeb", "9001")
}
