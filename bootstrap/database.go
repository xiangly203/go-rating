package bootstrap

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn = "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"

var Mysql *gorm.DB

func MysqlInit() {
	var err error
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
