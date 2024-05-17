package models

import (
	"github.com/leetcode-golang-classroom/gin-auth/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(uri string) {
	database, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		util.FailOnError(err, "Failed to connect to database")
	}
	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Blog{}, &User{})
}
