package handlers

import (
	countrydto "dumbmerch/dto/country"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type countryHandlers struct {
	CountryRepository repository.CountryRepository
}

func HandlerCountry(CountryRepository repository.CountryRepository) *countryHandlers {
	return &countryHandlers{CountryRepository}
}

func (h *countryHandlers) GetAllCountry(c echo.Context) error {
	countries, err := h.CountryRepository.GetAllCountry()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: countries})
}

func (h *countryHandlers) GetCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: country})

}

func (h *countryHandlers) CreateCountry(c echo.Context) error {
	request := new(countrydto.CreateCountryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	country := models.Country{
		Name: request.Name,
	}

	data, err := h.CountryRepository.CreateCountry(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)})
}

func (h *countryHandlers) UpdateCountry(c echo.Context) error {
	request := new(countrydto.UpdateCountryRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	coutry, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		coutry.Name = request.Name
	}

	data, err := h.CountryRepository.UpdateCountry(coutry)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry2(data)})
}

func (h *countryHandlers) DeleteCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CountryRepository.DeleteCountry(country, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry2(data)})
}

func convertResponseCountry(u models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func convertResponseCountry2(u models.Country) countrydto.CountryResponse2 {
	return countrydto.CountryResponse2{
		Name: u.Name,
	}
}
