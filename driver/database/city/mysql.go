package city

import (
	"context"
	"main-backend/bussiness/city"

	"gorm.io/gorm"
)

type cityRepository struct {
	conn *gorm.DB
}

func NewCityRepository(conn *gorm.DB) city.Repository {
	return &cityRepository{
		conn: conn,
	}
}

func (cr *cityRepository) Find(ctx context.Context) ([]city.Domain, error) {
	rec := []City{}
	err := cr.conn.Find(&rec).Error
	if err != nil {
		return []city.Domain{}, err
	}

	cityDomain := []city.Domain{}
	for _, value := range rec {
		cityDomain = append(cityDomain, value.ToDomain())
	}

	return cityDomain, nil
}

func (cr *cityRepository) FindByID(ctx context.Context, ID int) (city.Domain, error) {
	rec := City{}
	err := cr.conn.Find(&rec).Error
	if err != nil {
		return city.Domain{}, err
	}

	return rec.ToDomain(), nil
}
