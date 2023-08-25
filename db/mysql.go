package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-douyin/config"
)

var DB *gorm.DB

func init() {
	dsn := config.GetConfig().DatabaseURI
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic(err)
	}
	DB = db
}
