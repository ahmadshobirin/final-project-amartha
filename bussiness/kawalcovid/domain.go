package kawalcovid

import "context"

type Domain struct {
	Meninggal string `json:"meninggal"`
	Name      string `json:"name"`
	Positif   string `json:"positif"`
	Sembuh    string `json:"sembuh"`
}

type Usecase interface {
	GetAll(ctx context.Context) (Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) (Domain, error)
}
