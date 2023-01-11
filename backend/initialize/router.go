package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"timedeliver/middleware"
	"timedeliver/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors())

	deleverRouter := router.RouterGroupApp.Delever

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
		// 基础操作
		deleverRouter.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		//用户操作
		deleverRouter.InitUserRouter(PrivateGroup)

		//代办事项
		deleverRouter.InitTodingRouter(PrivateGroup)
	}

	return Router
}
