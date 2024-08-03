package main

import (
	"github.com/gin-gonic/gin"
	srv "liuche.com/project-common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "userWeb", "9001")
}
