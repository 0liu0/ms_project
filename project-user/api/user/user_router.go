package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (*UserRouter) Register(r *gin.Engine) {
	// 用户相关处理方法
	handler := New()

	// 路由
	g := r.Group("/project/login")
	{
		// 获取验证码
		g.POST("/getCaptcha", handler.GetCaptcha)
	}
}
