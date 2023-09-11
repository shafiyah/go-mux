package config

import (
	"github.com/shafiyah/go-mux/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go_api"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		models.Nationality{},
		models.Customer{},
		models.FamilyList{})
	DB = db
}
