package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"timedeliver/config"
)

var (
	Viper  *viper.Viper
	Config config.Server
	Zap    *zap.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)
