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

func (cr *queueRepository) FindByID(ctx context.Context, ID int) (queue.Domain, error) {
	rec := Queue{}
	err := cr.Conn.Preload("User").Preload("Clinic").Where("queues.id = ?", ID).Find(&rec).Error
	if err != nil {
		return queue.Domain{}, err
	}

	return rec.toDomain(), nil
}

func (cr *queueRepository) FindOne(ctx context.Context, userID, clinicID int, status string) (queue.Domain, error) {
	rec := Queue{}
	err := cr.Conn.Preload("User").Preload("Clinic").Where("user_id = ? AND clinic_id = ? AND status = ?", userID, clinicID, status).Find(&rec).Error
	if err != nil {
		return queue.Domain{}, err
	}

	return rec.toDomain(), nil
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
