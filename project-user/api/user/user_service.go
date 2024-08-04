package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	common "liuche.com/project-common"
	"liuche.com/project-common/constant"
	"liuche.com/project-common/logs"
	"liuche.com/project-common/tool"
	"liuche.com/project-user/pkg/dao"
	"liuche.com/project-user/pkg/model"
	"liuche.com/project-user/pkg/repo"
)

// 方法区
type UserHandler struct {
	cache repo.Cache
}

func New() *UserHandler {
	return &UserHandler{
		cache: dao.Rc,
	}
}

// GetCaptcha 获取手机验证码
func (handler *UserHandler) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	// 校验手机号码
	flag := tool.VerifyMobile(mobile)
	if !flag {
		logs.LOG.Error("mobile validate err!")
		ctx.JSON(http.StatusOK, result.Err(model.NoLegalMobile, "手机号校验失败！"))
		return
	}
	// 查看之前有没有发送过？有没有过期
	_, err := handler.cache.Get(constant.Registry_Key + mobile)
	if err != redis.Nil {
		logs.LOG.Warn("send code too busy!")
		ctx.JSON(http.StatusOK, result.Err(model.NoLegalMobile, "发送验证码太频繁！"))
		return
	}

	// 发送、存储
	code := "123456"
	go func() {
		time.Sleep(2 * time.Second)
		logs.LOG.Info("短信验证码发送成功！")
		// 存储至缓存服务器
		error := handler.cache.Put(constant.Registry_Key+mobile, code, 5*time.Minute)
		if error != nil {
			logs.LOG.Sugar().Errorf("Put code to redis err: 【%v】", error)
		} else {
			logs.LOG.Sugar().Infof("将手机号码和验证码存入Redis成功!key【%v】;code【%v】", constant.Registry_Key+mobile, code)
		}
	}()

	// 返回
	ctx.JSON(http.StatusOK, result.Ok("123456"))
}
