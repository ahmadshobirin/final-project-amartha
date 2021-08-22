package user

import (
	"context"
	"fmt"
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

func (uc *userUsecase) Fetch(ctx context.Context, roleCode string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userRepository.Fetch(ctx, roleCode, page, perpage)
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

	if email == "" {
		return Domain{}, messages.ErrInvalidParam
	}

	res, err := uc.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userUsecase) Store(ctx context.Context, data *Domain, roleID int) (res Domain, err error) {
	exist, _ := uc.userRepository.FindByEmail(ctx, data.Email)
	if exist.ID != 0 {
		return res, messages.ErrDataAlreadyExist
	}

	data.Password, _ = encrypt.Hash(data.Password)

	if roleID != 0 {
		data.RoleID = roleID
	}

	res, err = uc.userRepository.Store(ctx, data)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc *userUsecase) Update(ctx context.Context, data *Domain) (err error) {
	existedUsers, err := uc.userRepository.FindByID(ctx, data.ID)
	if err != nil {
		return err
	}

	fmt.Printf("exist user: %+v  \n", existedUsers)
	fmt.Println("after find by id")
	data.ID = existedUsers.ID

	if data.Password != "" {
		data.Password, _ = encrypt.Hash(data.Password)
	}

	err = uc.userRepository.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
