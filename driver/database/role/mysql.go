package role

import (
	"context"
	"main-backend/bussiness/role"

	"gorm.io/gorm"
)

type roleRepository struct {
	conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) role.Repository {
	return &roleRepository{
		conn: conn,
	}
}

func (cr *roleRepository) Find(ctx context.Context) ([]role.Domain, error) {
	rec := []Role{}
	err := cr.conn.Find(&rec).Error
	if err != nil {
		return []role.Domain{}, err
	}

	roleDomain := []role.Domain{}
	for _, value := range rec {
		roleDomain = append(roleDomain, value.ToDomain())
	}

	return roleDomain, nil
}

func (cr *roleRepository) FindByID(ctx context.Context, ID int) (role.Domain, error) {
	rec := Role{}
	err := cr.conn.Find(&rec).Error
	if err != nil {
		return role.Domain{}, err
	}

	return rec.ToDomain(), nil
}
