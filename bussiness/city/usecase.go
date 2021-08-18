package city

import (
	"context"
	"main-backend/helper/messages"
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

func (uc *CityUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := uc.cityRespository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (uc *CityUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := uc.cityRespository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *CityUsecase) GetByName(ctx context.Context, name string) (Domain, error) {
	resp, err := uc.cityRespository.GetByName(ctx, name)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}

func (uc *CityUsecase) Store(ctx context.Context, data *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, _ := uc.cityRespository.GetByName(ctx, data.Name)
	if exist != (Domain{}) {
		return messages.ErrDuplicateData
	}

	err := uc.cityRespository.Store(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CityUsecase) Update(ctx context.Context, data *Domain) (err error) {
	exist, err := uc.cityRespository.FindByID(ctx, data.ID)
	if err != nil {
		return err
	}
	data.ID = exist.ID

	err = uc.cityRespository.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CityUsecase) Delete(ctx context.Context, data *Domain) (err error) {
	existedWebinar, err := uc.cityRespository.FindByID(ctx, data.ID)
	if err != nil {
		return err
	}
	data.ID = existedWebinar.ID

	err = uc.cityRespository.Delete(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
