package models

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	Id        uint           `json:"id" gorm:"primarykey" `
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index" `
}
