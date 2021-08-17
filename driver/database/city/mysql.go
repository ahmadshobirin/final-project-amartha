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

func (repo *cityRepository) GetAll(ctx context.Context) ([]city.Domain, error) {
	rec := []City{}
	err := repo.conn.Find(&rec).Error
	if err != nil {
		return []city.Domain{}, err
	}

	cityDomain := []city.Domain{}
	for _, value := range rec {
		cityDomain = append(cityDomain, value.ToDomain())
	}

	return cityDomain, nil
}

func (repo *cityRepository) FindByID(ctx context.Context, ID int) (city.Domain, error) {
	rec := City{}
	err := repo.conn.Find(&rec).Error
	if err != nil {
		return city.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (repo *cityRepository) GetByName(ctx context.Context, name string) (city.Domain, error) {
	rec := City{}
	err := repo.conn.Where("name = ?", name).Find(&rec).Error
	if err != nil {
		return city.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (repo cityRepository) Store(ctx context.Context, data *city.Domain) (err error) {
	rec := fromDomain(data)
	result := repo.conn.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return err
}

func (repo *cityRepository) Update(ctx context.Context, data *city.Domain) (err error) {
	rec := fromDomain(data)

	result := repo.conn.Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *cityRepository) Delete(ctx context.Context, data *city.Domain) (err error) {
	rec := fromDomain(data)

	result := repo.conn.Delete(rec)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
