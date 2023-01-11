package config

type System struct {
	Addr     int    `mapstructure:"addr" json:"addr" yaml:"addr"`                // 端口值
	DbType   string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`       // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseRedis bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
}
