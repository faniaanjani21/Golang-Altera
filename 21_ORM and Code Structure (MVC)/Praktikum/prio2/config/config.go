package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"prio2/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigration()
}

func initMigration() {
	DB.AutoMigrate(&models.Buku{})
	DB.AutoMigrate(&models.User{})
}
