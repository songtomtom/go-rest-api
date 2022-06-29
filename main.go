package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:qweqweqwe@tcp(127.0.0.1:3306)/go_rest_api"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

}
