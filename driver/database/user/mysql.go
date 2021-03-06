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

func (repo *mysqlUsersRepository) Fetch(ctx context.Context, roleCode string, page, perpage int) ([]user.Domain, int, error) {
	rec := []User{}

	offset := (page - 1) * perpage
	query := repo.Conn.Preload("Role").Distinct().Joins("JOIN roles ON roles.id = users.role_id")
	if roleCode != "" {
		query = query.Where("roles.code = ?", roleCode)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []user.Domain{}, 0, err
	}

	var totalData int64
	err = query.Count(&totalData).Error
	if err != nil {
		return []user.Domain{}, 0, err
	}

	var domainNews []user.Domain
	for _, value := range rec {
		domainNews = append(domainNews, *value.ToDomain())
	}
	return domainNews, int(totalData), nil
}

func (repo *mysqlUsersRepository) FindByID(ctx context.Context, userID int) (user.Domain, error) {
	rec := User{}
	err := repo.Conn.Preload("Role").Where("users.id = ?", userID).First(&rec).Error
	if err != nil {
		return user.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repo *mysqlUsersRepository) FindByEmail(ctx context.Context, email string) (user.Domain, error) {
	rec := User{}
	err := repo.Conn.Preload("Role").Where("email = ?", email).First(&rec).Error
	if err != nil {
		return user.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repo *mysqlUsersRepository) Store(ctx context.Context, data *user.Domain) (res user.Domain, err error) {
	rec := fromDomain(data)
	result := repo.Conn.Create(&rec)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	err = repo.Conn.Preload("Role").First(&rec, rec.ID).Error
	if err != nil {
		return user.Domain{}, result.Error
	}

	return *rec.ToDomain(), err
}

func (repo *mysqlUsersRepository) Update(ctx context.Context, data *user.Domain) error {
	rec := fromDomain(data)
	result := repo.Conn.Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
