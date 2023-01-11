package deliver

import "timedeliver/service"

type ApiGroup struct {
	UserApi
	TodingApi
}

var (
	userService   = service.ServiceGroupApp.DeliverServiceGroup.UserService
	utilsService  = service.ServiceGroupApp.DeliverServiceGroup.UtilsService
	jwtService    = service.ServiceGroupApp.DeliverServiceGroup.JwtService
	todingService = service.ServiceGroupApp.DeliverServiceGroup.TodoingService
)
