package deliver

import (
	"github.com/gin-gonic/gin"
	v1 "timedeliver/api/v1"
)

type TodingRouter struct {
}

func (e *TodingRouter) InitTodingRouter(Router *gin.RouterGroup) {
	todingRouter := Router.Group("todo")
	todoApi := v1.ApiGroupApp.DeliverApiGroup.TodingApi
	{
		todingRouter.POST("create", todoApi.CreateToding)        // 创建代办事项
		todingRouter.POST("monthtodo", todoApi.GetTodingByMonth) //根据年月获得当月的todo

	}
}
