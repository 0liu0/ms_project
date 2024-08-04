package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"liuche.com/project-user/api/user"
)

// 路由实现接口
type Router interface {
	Register(r *gin.Engine)
}

// 路由注册器
type RegisterRouter struct {
}

// 新建路由
func NewRouter() RegisterRouter {
	return RegisterRouter{}
}

// 实现路由方法
func (*RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Register(r)
}

// 初始化路由
func InitRouter(r *gin.Engine) {
	fmt.Println("我开始注册了哦！！！")

	router := NewRouter()
	// 新建user路由
	router.Route(&user.UserRouter{}, r)
}
