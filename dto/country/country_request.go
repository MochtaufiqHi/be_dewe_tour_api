package countrydto

type CreateCountryRequest struct {
	Name string `json:"name" from:"name" validate:"required"`
}

type UpdateCountryRequest struct {
	Name string `json:"name" from:"name"`
}
