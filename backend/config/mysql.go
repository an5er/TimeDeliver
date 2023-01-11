package config

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`                        // 服务器地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`                        //端口
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`               // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"`            // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"`            // 数据库密码
	Engine   string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"` //数据库引擎，默认InnoDB
	Config   string `mapstructure:"db_config" json:"db_config" yaml:"db_config"`         // 高级配置
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
