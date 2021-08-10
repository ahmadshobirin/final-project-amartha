package city

import (
	"context"
	"time"
)

type CityUsecase struct {
	cityRespository Repository
	contextTimeout  time.Duration
}

func NewCityUsecase(timeout time.Duration, cr Repository) Usecase {
	return &CityUsecase{
		contextTimeout:  timeout,
		cityRespository: cr,
	}
}

func (cu *CityUsecase) Find(ctx context.Context) ([]Domain, error) {
	resp, err := cu.cityRespository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (cu *CityUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	return Domain{}, nil
}
