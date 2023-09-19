package db

import (
	"go-web-service/pkg/config"
	"go-web-service/pkg/repo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 暴露获取连接的方法，其他留给上层模块实现
var _db *gorm.DB

func InitDb(dbConfig *config.DbConfig) error {
	dbInstance, err := gorm.Open(mysql.Open(dbConfig.Dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	_db = dbInstance
	repo.InitRepos(_db)
	return nil
}

func GetDb() *gorm.DB {
	return _db
}
