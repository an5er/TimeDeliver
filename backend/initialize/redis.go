package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"timedeliver/global"
)

func Redis() {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Zap.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.Zap.Info("redis connect ping response:", zap.String("pong", pong))
		global.Redis = client
	}
}
