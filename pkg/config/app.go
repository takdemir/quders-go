package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DB_URL")))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
