package model

import (
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, &Task{})
	return db
}
