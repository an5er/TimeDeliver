package deliver

import (
	"timedeliver/global"
)

type UtilsService struct{}

var UtilsServiceApp = new(UtilsService)

func (utilsService *UtilsService) GetUidByUname(uname string) int {
	var uid int
	global.DB.Table("users").Select("uid").Where("uname = ?", uname).Find(&uid)
	return uid
}
