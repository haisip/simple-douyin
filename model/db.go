package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-douyin/config"
)

var DB *gorm.DB
var dsn string

func init() {
	dsn = config.GetConfig().DatabaseURL
}

func InitDB() error {
	//db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})  // sqlite
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // mysql
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&User{}, &Comment{}, &Video{},
		&UserUser{}, &UserVideo{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
