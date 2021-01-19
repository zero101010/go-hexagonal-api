package database

import (
	"go-api-hexagonal/domain"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBaseMysql() {
	dsn := "root:12345678@tcp(35.188.64.151:3306)/go-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&domain.User{})

	DB = db
}

func ConnectDataBaseSqlight() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&domain.User{})

	DB = db

}
