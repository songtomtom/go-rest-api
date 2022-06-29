package app

import (
	"songtomtom/rest_api/app/model"
	"songtomtom/rest_api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) Init() {

	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	model.Migration(db)
}
