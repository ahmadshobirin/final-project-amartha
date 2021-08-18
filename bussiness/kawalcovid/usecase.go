package kawalcovid

import (
	"context"
	"time"
)

type kawalCovidUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewKawalCovidUsecase(timeout time.Duration, wr Repository) Usecase {
	return &kawalCovidUsecase{
		contextTimeout: timeout,
		repo:           wr,
	}
}

func (uc *kawalCovidUsecase) GetAll(ctx context.Context) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.repo.GetAll(ctx)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
