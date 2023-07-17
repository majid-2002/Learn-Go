package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	passDb := os.Getenv("DB_PASS")

	dsn := "abdul_majid:" + passDb + "@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
