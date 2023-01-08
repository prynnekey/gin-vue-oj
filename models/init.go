package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = Init()

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
