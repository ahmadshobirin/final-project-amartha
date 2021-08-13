package role

import (
	"context"
	"time"
)

type RoleUsecase struct {
	roleRepository Repository
	contextTimeout time.Duration
}

func NewRoleUsecase(timeout time.Duration, cr Repository) Usecase {
	return &RoleUsecase{
		contextTimeout: timeout,
		roleRepository: cr,
	}
}

func (cu *RoleUsecase) Find(ctx context.Context) ([]Domain, error) {
	resp, err := cu.roleRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}

func (cu *RoleUsecase) FindByID(ctx context.Context, id int) (Domain, error) {
	resp, err := cu.roleRepository.FindByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}
