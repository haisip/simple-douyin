package model

import (
	"fmt"
	"gorm.io/gorm"
)
import dbPackage "simple-douyin/db"

var db *gorm.DB

func init() {
	db = dbPackage.DB
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(&User{}, &Comment{}, &Video{}, &UserUser{}, &UserVideo{})
	if err != nil {
		panic(err)
	}
}

func CreateTables() error {
	fmt.Println(db)
	return db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(&User{}, &Comment{}, &Video{}, &UserUser{}, &UserVideo{})
}
