package user

import (
	"context"
	"database/sql"
	"main-backend/helper/encrypt"
	"main-backend/helper/messages"
	"time"
)

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
}

func NewUserUsecase(timeout time.Duration, ur Repository) Usecase {
	return &userUsecase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *userUsecase) FindByID(ctx context.Context, ID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if ID <= 0 {
		return Domain{}, messages.ErrIDNotFound
	}

	res, err := uc.userRepository.FindByID(ctx, ID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userUsecase) FindByEmail(ctx context.Context, email string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userUsecase) Store(ctx context.Context, data *Domain) (err error) {
	exist, _ := uc.FindByEmail(ctx, data.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if exist != (Domain{}) {
		return messages.ErrDataAlreadyExist
	}

	data.Password, err = encrypt.Hash(data.Password)
	if err != nil {
		return messages.ErrInternalServer
	}

	err = uc.userRepository.Store(ctx, data)
	if err != nil {
		return err
	}

	return err
}
