package database

import (
	"go-admin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:myroot@/world"), &gorm.Config{})
	if err != nil {
		panic("could not connection to the database")
	}
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{})

}
