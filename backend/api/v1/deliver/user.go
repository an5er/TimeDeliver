package deliver

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"timedeliver/global"
	"timedeliver/model/request"
	"timedeliver/model/response"
	"timedeliver/util"
)

type UserApi struct{}

func (e *UserApi) CreateUser(c *gin.Context) {
	var user request.User

	err := c.ShouldBindJSON(&user)

	if util.CheckNull(user) != true {
		response.FailWithMessage("参数不完整", c)
	}

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.CreateUser(user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.Zap.Info("创建用户发生错误: " + err.Error())
		return
	}
	global.Zap.Info("成功创建用户: " + user.Uname)
	response.OkWithMessage("用户创建成功", c)
}

func (e *UserApi) LoginUser(c *gin.Context) {
	var user request.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.LoginUser(user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	e.TokenNext(user, c)

}

func (e *UserApi) TokenNext(user request.User, c *gin.Context) string {

	uid := utilsService.GetUidByUname(user.Uname)

	j := &util.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UID:      uid,
		Username: user.Uname,
	})
	global.Zap.Info("签发令牌给 ID: " + string(uid) + "name: " + user.Uname)
	token, err := j.CreateToken(claims)

	if err != nil {
		global.Zap.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return ""
	}

	if _, err := jwtService.GetRedisJWT(user.Uname); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Uname); err != nil {
			global.Zap.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return ""
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.Zap.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {

		if err := jwtService.SetRedisJWT(token, user.Uname); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return ""
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)

	}
	return ""
}

func (e *UserApi) ChangePasswd(c *gin.Context) {
	var newUser request.NewUser

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if util.CheckNull(newUser) != true {
		response.FailWithMessage("参数不完整", c)
		return
	}

	err = userService.ChangePasswd(newUser)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.Zap.Info("创建用户发生错误: " + err.Error())
		return
	}
	global.Zap.Info("用户成功更改密码: " + newUser.User.Uname)
	jwtService.GetRedisJWT(newUser.User.Uname)

	var user request.User
	user.Uname = newUser.User.Uname
	e.TokenNext(user, c)

	response.OkWithMessage("密码更改成功", c)
}
