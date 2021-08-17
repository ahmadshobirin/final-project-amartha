package user

import (
	"context"
	user "main-backend/bussiness/user"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) user.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (nr *mysqlUsersRepository) Fetch(ctx context.Context, page, perpage int) ([]user.Domain, int, error) {
	rec := []User{}

	offset := (page - 1) * perpage
	err := nr.Conn.Preload("Roles").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []user.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []user.Domain{}, 0, err
	}

	var domainNews []user.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (nr *mysqlUsersRepository) FindByID(ctx context.Context, newsId int) (user.Domain, error) {
	rec := User{}
	err := nr.Conn.Where("id = ?", newsId).First(&rec).Error
	if err != nil {
		return user.Domain{}, err
	}

	return rec.toDomain(), nil
}

func (nr *mysqlUsersRepository) FindByEmail(ctx context.Context, email string) (user.Domain, error) {
	rec := User{}
	err := nr.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return user.Domain{}, err
	}

	return rec.toDomain(), nil
}

func (nr *mysqlUsersRepository) Store(ctx context.Context, data *user.Domain) (res user.Domain, err error) {
	rec := fromDomain(data)
	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	err = nr.Conn.Preload("Role").First(&rec, rec.ID).Error
	if err != nil {
		return user.Domain{}, result.Error
	}

	return rec.toDomain(), err
}
