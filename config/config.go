package config

import "fmt"

type DB struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

func GetDSN() string {
	db := &DB{
		Username: "root",
		Password: "qweqweqwe",
		Host:     "127.0.0.1",
		Port:     3306,
		Name:     "go_rest_api",
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
