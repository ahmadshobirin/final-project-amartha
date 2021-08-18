package queue

import (
	"context"
	"main-backend/helper/messages"
	"time"
)

type queueUsecase struct {
	queueRepository Repository
	contextTimeout  time.Duration
}

func NewQueueUsecase(timeout time.Duration, wr Repository) Usecase {
	return &queueUsecase{
		contextTimeout:  timeout,
		queueRepository: wr,
	}
}

func (uc *queueUsecase) FindOne(ctx context.Context, userID, clinicID int, status string) (res Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.queueRepository.FindOne(ctx, userID, clinicID, status)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc *queueUsecase) Store(ctx context.Context, userID int, data *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	data.UserID = userID
	data.Status = StatusPending

	exist, _ := uc.queueRepository.FindOne(ctx, data.UserID, data.ClinicID, StatusPending)
	if exist != (Domain{}) {
		return messages.ErrDuplicateData
	}

	err := uc.queueRepository.Store(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *queueUsecase) FindByID(ctx context.Context, ID int) (res Domain, err error) {
	res, err = uc.queueRepository.FindByID(ctx, ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc *queueUsecase) Update(ctx context.Context, data *Domain) error {
	exist, err := uc.queueRepository.FindByID(ctx, data.ID)
	if err != nil {
		return err
	}

	data.ID = exist.ID
	data.Status = StatusPaid

	err = uc.queueRepository.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
