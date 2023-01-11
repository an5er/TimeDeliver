package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"timedeliver/global"
)

func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		//todo
		return nil
	default:
		return GormMysql()
	}
}
func GormMysql() *gorm.DB {
	m := global.Config.Mysql

	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err != nil {
		fmt.Println("Open datatbase failed!")
		fmt.Println(m.Dsn())
	}
	db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetMaxIdleConns(200)
	return db
}
