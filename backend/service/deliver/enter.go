package deliver

import (
	"context"
	"timedeliver/global"
	"timedeliver/util"
)

type ServiceGroup struct {
	UserService
	UtilsService
	JwtService
	TodoingService
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.Redis.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := util.ParseDuration(global.Config.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
