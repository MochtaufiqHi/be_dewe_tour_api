package countrydto

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CountryResponse2 struct {
	Name string `json:"name"`
}
