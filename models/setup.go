package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/jwt_go"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db

}
