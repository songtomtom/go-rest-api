package model

import (
	"fmt"
	"songtomtom/rest_api/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() *gorm.DB {
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connect error.")
	}
	db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.Create(&user)
	fmt.Println(user)
	return db
}
