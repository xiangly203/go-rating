package dal

import (
	"go_gin/biz/dal/cache"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlUrl = "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"

var Mysql *gorm.DB

func MysqlInit() {
	var err error
	Mysql, err = gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

var postgresUrl = "postgres:123456@tcp(localhost:5432)/test?charset=utf8&parseTime=True&loc=Local"

var Postgres *gorm.DB

func PostgresInit() {
	var err error
	Postgres, err = gorm.Open(postgres.Open(postgresUrl), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

func CacheInit() {
	cache.NewRedisCache(cache.RedisConf{
		Addr:     "localhost:6379",
		Password: "",
		Db:       0,
	})
}
