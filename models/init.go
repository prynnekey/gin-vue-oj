package models

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = Init()
var REDIS = initRedis()

func Init() *gorm.DB {
	dsn := "root:prynnekey@tcp(127.0.0.1:3306)/gin-vue-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 开启日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("connect to database failed:" + err.Error())
	}

	return db
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.138.128:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	return rdb
}
