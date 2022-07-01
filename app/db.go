package app

import (
	"songtomtom/rest_api/app/model"
	"songtomtom/rest_api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() *gorm.DB {
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connect error.")
	}
	db.AutoMigrate(&model.User{})
	return db
}
