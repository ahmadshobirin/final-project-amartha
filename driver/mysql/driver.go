package mysql_driver

import (
	"fmt"
	"log"
	"main-backend/driver/database/city"
	"main-backend/driver/database/clinic"
	"main-backend/driver/database/queue"
	"main-backend/driver/database/role"
	"main-backend/driver/database/user"
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
	Seeder(db)

	return db
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&city.City{}, &role.Role{}, &user.User{}, &clinic.Clinic{}, &queue.Queue{})
}

func Seeder(db *gorm.DB) {
	cities := []city.City{}
	db.Find(&cities)
	if len(cities) == 0 {
		cities = []city.City{
			{Code: "SBY", Name: "Surabaya", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Code: "SDA", Name: "Sidoarjo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Code: "MJK", Name: "Mojokerto", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Code: "JKT", Name: "Jakarta", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		db.Create(&cities)
	}

	var roles = []role.Role{}
	db.Find(&roles)
	if len(roles) == 0 {
		roles = []role.Role{
			{Code: "SA", Name: "Superadmin", Status: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Code: "AM", Name: "Admin", Status: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Code: "US", Name: "User", Status: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		db.Create(&roles)
	}
}
