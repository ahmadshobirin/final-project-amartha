package kawalcovid

type ResponseKawalKovid []struct {
	Meninggal string `json:"meninggal"`
	Name      string `json:"name"`
	Positif   string `json:"positif"`
	Sembuh    string `json:"sembuh"`
}
