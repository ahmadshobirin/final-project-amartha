package clinic

import (
	"main-backend/driver/database/city"
	"main-backend/driver/database/user"
	"time"
)

type Clinic struct {
	ID          int `gorm:"primaryKey"`
	UserID      int
	User        user.User
	CityID      int
	City        city.City
	Name        string
	Description string
	OpenTime    string
	CloseTime   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
