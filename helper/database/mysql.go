package database

import (
	"fmt"
	"log"
	cityRepo "main-backend/driver/database/city"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)
	// Seeder(db)

	return db
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&cityRepo.City{})
}

func Seeder(db *gorm.DB) {
	var cities = []cityRepo.City{
		{Code: "SBY", Name: "Surabaya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Code: "SDA", Name: "Sidoarjo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Code: "MJK", Name: "Mojokerto", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Code: "JKT", Name: "Jakarta", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	db.Create(&cities)
}
