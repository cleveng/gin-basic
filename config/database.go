package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)
// DB gorm connector
var DB *gorm.DB


// ConnectDB connect to db
func ConnectDB() {
	var err error
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config("DB_USERNAME"),
		Config("DB_PASSWORD"),
		Config("DB_HOST"),
		Config("DB_PORT"),
		Config("DB_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),	//追踪日志
	})
	if err != nil {
		panic("数据库链接失败，疑似开发环境未启动")
	}
	if sqlDB, err := DB.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)  //SetMaxIdleConns用于设置闲置的连接数
		sqlDB.SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
		sqlDB.SetConnMaxLifetime(time.Hour)	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	}
	// 启用Logger，显示详细日志
	if Config("APP_ENV") == "development" {
		DB = DB.Debug()
	}
}
