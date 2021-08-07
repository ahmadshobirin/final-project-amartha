package config

import (
	"fmt"
	"main-backend/app/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	DBConnection string
	DBName       string
	DBHost       string
	DBPort       string
	DBUsername   string
	DBPassword   string
}

func connectionMap() config {
	conf := config{
		DBConnection: os.Getenv("DB_CONNECTION"),
		DBName:       os.Getenv("DB_NAME"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
	}

	return conf
}

func assembleConfig() string {
	conf := connectionMap().DBUsername + ":" +
		connectionMap().DBPassword + "@(" +
		connectionMap().DBHost + ":" +
		connectionMap().DBPort + ")/" +
		connectionMap().DBName + "?" +
		"parseTime=true"

	// fmt.Println("conf_DB", conf)
	return conf
}

func InitDB() *gorm.DB {
	var err error
	initDB, err := gorm.Open(mysql.Open(assembleConfig()), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed Connection Database")
	}

	return initDB
}

func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(&models.City{})
}
