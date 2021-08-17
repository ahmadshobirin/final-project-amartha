package clinic

import (
	"context"
	"main-backend/helper/messages"
	"time"
)

type clinicUsecase struct {
	clinicRepository Repository
	contextTimeout   time.Duration
}

func NewClinicUsecase(timeout time.Duration, wr Repository) Usecase {
	return &clinicUsecase{
		contextTimeout:   timeout,
		clinicRepository: wr,
	}
}

func (uc *clinicUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.clinicRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *clinicUsecase) GetByID(ctx context.Context, clinicID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if clinicID <= 0 {
		return Domain{}, messages.ErrNotFound
	}
	res, err := uc.clinicRepository.GetByID(ctx, clinicID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *clinicUsecase) GetByUserID(ctx context.Context, userID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if userID <= 0 {
		return Domain{}, messages.ErrNotFound
	}
	res, err := uc.clinicRepository.GetByUserID(ctx, userID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *clinicUsecase) Store(ctx context.Context, data *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, _ := uc.clinicRepository.GetByUserID(ctx, data.UserID)
	if exist != (Domain{}) {
		return messages.ErrDuplicateData
	}

	err := uc.clinicRepository.Store(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *clinicUsecase) Update(ctx context.Context, data *Domain) (err error) {
	exist, err := uc.clinicRepository.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}
	data.ID = exist.ID

	err = uc.clinicRepository.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *clinicUsecase) Delete(ctx context.Context, data *Domain) (err error) {
	exist, err := uc.clinicRepository.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}
	data.ID = exist.ID

	err = uc.clinicRepository.Delete(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
