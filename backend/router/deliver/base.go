package deliver

import (
	"github.com/gin-gonic/gin"
	v1 "timedeliver/api/v1"
)

type BaseRouter struct {
}

func (e *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.DeliverApiGroup.UserApi
	{
		baseRouter.POST("register", baseApi.CreateUser) // 用户注册
		baseRouter.POST("login", baseApi.LoginUser)     // 用户登录
	}
}
