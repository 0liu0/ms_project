package main

import (
	"log"

	"github.com/gin-gonic/gin"
	srv "liuche.com/project-common"
	"liuche.com/project-common/logs"
	"liuche.com/project-user/router"
)

func main() {
	r := gin.Default()
	// 初始化路由
	router.InitRouter(r)
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "..//logs//debug//project-debug.log",
		InfoFileName:  "..//logs//info//project-info.log",
		WarnFileName:  "..//logs//error//project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	srv.Run(r, "userWeb", "9001")
}
