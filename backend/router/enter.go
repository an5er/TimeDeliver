package router

import "timedeliver/router/deliver"

type RouterGroup struct {
	Delever deliver.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
