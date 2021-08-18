package queue

import (
	"context"
	"main-backend/bussiness/queue"

	"gorm.io/gorm"
)

type queueRepository struct {
	Conn *gorm.DB
}

func NewQueueRepository(conn *gorm.DB) queue.Repository {
	return &queueRepository{
		Conn: conn,
	}
}

func (repo *queueRepository) Fetch(ctx context.Context, userID, page, perpage int) ([]queue.Domain, int, error) {
	rec := []Queue{}

	offset := (page - 1) * perpage
	err := repo.Conn.Preload("User").Preload("Clinic").Where("queues.id = ?", userID).Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []queue.Domain{}, 0, err
	}

	var totalData int64
	err = repo.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []queue.Domain{}, 0, err
	}

	var data []queue.Domain
	for _, value := range rec {
		data = append(data, *value.ToDomain())
	}
	return data, int(totalData), nil
}

func (cr *queueRepository) FindByID(ctx context.Context, ID int) (queue.Domain, error) {
	rec := Queue{}
	err := cr.Conn.Preload("User").Preload("Clinic").Preload("Clinic.User").Where("queues.id = ?", ID).Find(&rec).Error
	if err != nil {
		return queue.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (cr *queueRepository) FindOne(ctx context.Context, userID, clinicID int, status string) (queue.Domain, error) {
	rec := Queue{}
	err := cr.Conn.Preload("User").Preload("Clinic").Where("user_id = ? AND clinic_id = ? AND status = ?", userID, clinicID, status).Find(&rec).Error
	if err != nil {
		return queue.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repo *queueRepository) Store(ctx context.Context, data *queue.Domain) (err error) {
	rec := fromDomain(data)

	result := repo.Conn.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return err
}

func (repo *queueRepository) Update(ctx context.Context, data *queue.Domain) (err error) {
	rec := fromDomain(data)

	result := repo.Conn.Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
