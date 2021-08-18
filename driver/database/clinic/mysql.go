package clinic

import (
	"context"
	"main-backend/bussiness/clinic"

	"gorm.io/gorm"
)

type mysqlClinicRepository struct {
	Conn *gorm.DB
}

func NewClinicRepository(conn *gorm.DB) clinic.Repository {
	return &mysqlClinicRepository{
		Conn: conn,
	}
}

func (repo *mysqlClinicRepository) Fetch(ctx context.Context, page, perpage int) ([]clinic.Domain, int, error) {
	rec := []Clinic{}

	offset := (page - 1) * perpage
	err := repo.Conn.Preload("User").Preload("City").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []clinic.Domain{}, 0, err
	}

	var totalData int64
	err = repo.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []clinic.Domain{}, 0, err
	}

	var domainNews []clinic.Domain
	for _, value := range rec {
		domainNews = append(domainNews, *value.ToDomain())
	}
	return domainNews, int(totalData), nil
}

func (repo *mysqlClinicRepository) GetByID(ctx context.Context, ID int) (res clinic.Domain, err error) {
	rec := Clinic{}
	err = repo.Conn.Preload("User").Preload("City").Where("clinics.id = ?", ID).First(&rec).Error
	if err != nil {
		return clinic.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repo *mysqlClinicRepository) GetByUserID(ctx context.Context, userID int) (res clinic.Domain, err error) {
	rec := Clinic{}
	err = repo.Conn.Preload("User").Preload("City").Where("clinics.user_id = ?", userID).First(&rec).Error
	if err != nil {
		return clinic.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repo mysqlClinicRepository) Store(ctx context.Context, data *clinic.Domain) (err error) {
	rec := fromDomain(data)
	result := repo.Conn.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return err
}

func (repo *mysqlClinicRepository) Update(ctx context.Context, data *clinic.Domain) (err error) {
	rec := fromDomain(data)

	result := repo.Conn.Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *mysqlClinicRepository) Delete(ctx context.Context, data *clinic.Domain) (err error) {
	rec := fromDomain(data)

	result := repo.Conn.Delete(rec)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
