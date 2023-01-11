package deliver

import (
	"errors"
	"gorm.io/gorm"
	"timedeliver/global"
	"timedeliver/model"
	"timedeliver/model/request"
	"timedeliver/util"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (userService *UserService) CreateUser(user request.User) (err error) {
	if !errors.Is(global.DB.Where("uname = ?", user.Uname).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}

	var cUser model.Users
	cUser.Uname = user.Uname
	cUser.Passwd = string(util.BcryptE(user.Passwd))

	return global.DB.Create(&cUser).Error
}

func (userService *UserService) LoginUser(user request.User) (err error) {
	var cUser model.Users

	if errors.Is(global.DB.Where("uname = ?", user.Uname).First(&cUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名或密码错误")
	}

	if !util.BcryptC(cUser.Passwd, user.Passwd) {
		return errors.New("用户名或密码错误")
	}

	return nil
}

func (userService *UserService) ChangePasswd(user request.NewUser) (err error) {
	var cUser model.Users

	if errors.Is(global.DB.Where("uname = ?", user.User.Uname).First(&cUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名或密码错误")
	}

	if !util.BcryptC(cUser.Passwd, user.User.Passwd) {
		return errors.New("用户名或密码错误")
	}

	cUser.Passwd = string(util.BcryptE(user.NewPasswd))
	global.Zap.Info("uname:" + user.User.Uname + "更改了密码")

	return global.DB.Where("uname = ?", user.User.Uname).Save(&cUser).Error
}
