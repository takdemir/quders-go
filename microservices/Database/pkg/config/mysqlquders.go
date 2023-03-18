package config

import (
	"database/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
}

func Connect() *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: os.Getenv("MYSQL_DB_URL")}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Currency{},
		&models.Company{},
		&models.User{},
	)
	if err != nil {
		return
	}
}
