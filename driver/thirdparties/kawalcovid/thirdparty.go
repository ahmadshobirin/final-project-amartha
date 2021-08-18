package kawalcovid

import (
	"main-backend/bussiness/kawalcovid"

	"context"
	"encoding/json"
	"net/http"
)

type Kawalcovid struct {
	httpClient http.Client
}

func NewKawalcovid() kawalcovid.Repository {
	return &Kawalcovid{
		httpClient: http.Client{},
	}
}

func (repo *Kawalcovid) GetAll(ctx context.Context) (res kawalcovid.Domain, err error) {
	req, _ := http.NewRequest("GET", "https://api.kawalcorona.com/indonesia", nil)
	resp, err := repo.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	bind := ResponseKawalKovid{}
	err = json.NewDecoder(resp.Body).Decode(&bind)
	if err != nil {
		return res, err
	}

	res = kawalcovid.Domain{}
	for _, data := range bind {
		res.Meninggal = data.Meninggal
		res.Sembuh = data.Sembuh
		res.Positif = data.Sembuh
		res.Name = data.Name
	}

	return res, nil
}
