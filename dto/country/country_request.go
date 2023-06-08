package countrydto

type CreateCountryRequest struct {
	Name string `json:"name" from:"name"`
}

type UpdateCountryRequest struct {
	Name string `json:"name" from:"name"`
}
