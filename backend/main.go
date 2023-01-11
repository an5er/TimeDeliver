package main

import (
	"fmt"
	"go.uber.org/zap"
	"timedeliver/core"
	"timedeliver/global"
	"timedeliver/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.Viper = core.Viper()

	global.Zap = core.Zap()
	zap.ReplaceGlobals(global.Zap)
	defer global.Zap.Sync()

	global.DB = initialize.Gorm()
	Router := initialize.Routers()

	initialize.Redis()

	address := fmt.Sprintf("0.0.0.0:%d", global.Config.System.Addr)
	err := Router.Run(address)

	switch err {
	case nil:
		global.Zap.Info("服务开始运行")
	default:
		panic(err)
	}
}
