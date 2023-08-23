package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-douyin/config"
)

var DB *gorm.DB
var dsn string

func init() {
	dsn = config.GetConfig().DatabaseURI
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, err
	}

	// 设置默认表选项，包括默认字符集和默认引擎
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(&User{}, &Comment{}, &Video{}, &UserUser{}, &UserVideo{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
