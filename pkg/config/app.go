package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
}

func Connect() {
	d, err := gorm.Open(mysql.New(mysql.Config{DSN: os.Getenv("MYSQL_DB_URL")}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
