package deliver

import (
	"github.com/gin-gonic/gin"
	v1 "timedeliver/api/v1"
)

type UserRouter struct {
}

func (e *BaseRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.DeliverApiGroup.UserApi
	{
		userRouter.POST("edit", userApi.ChangePasswd) // 用户改密码

	}
}
