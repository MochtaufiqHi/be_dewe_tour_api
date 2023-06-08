package handlers

import (
	dto "dumbmerch/dto/result"
	tripdto "dumbmerch/dto/trip"
	"dumbmerch/models"
	"dumbmerch/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type tripHandlers struct {
	TripRepository repository.TripRepository
}

func HandlerTrip(TripRepository repository.TripRepository) *tripHandlers {
	return &tripHandlers{TripRepository}
}

var path_file = "http://localhost:5000/uploads/"

func (h *tripHandlers) GetAllTrip(c echo.Context) error {
	trips, err := h.TripRepository.GetAllTrip()
	// fmt.Println(trips)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	for i, p := range trips {
		trips[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: trips})
}

func (h *tripHandlers) GetTrip(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(trip)})
}

func (h *tripHandlers) CreateTrip(c echo.Context) error {
	// if you want to use request body raw json
	// request := new(tripdto.CreateTripRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// if you want to use request body form-data
	dataFile := c.Get("dataFile").(string)
	fmt.Println("This is data file", dataFile)

	day, _ := strconv.Atoi(c.FormValue("day"))
	night, _ := strconv.Atoi(c.FormValue("night"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	countryId, _ := strconv.Atoi(c.FormValue("country_id"))
	quota, _ := strconv.Atoi(c.FormValue("quota"))

	request := tripdto.CreateTripRequest{
		Title:          c.FormValue("title"),
		CountryID:      countryId,
		Acommodation:   c.FormValue("acommodation"),
		Transportation: c.FormValue("transportation"),
		Eat:            c.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       c.FormValue("dateTrip"),
		Price:          price,
		Quota:          quota,
		Description:    c.FormValue("description"),
		Image:          dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	idCountry, _ := h.TripRepository.GetCountryByID(request.CountryID)
	// fmt.Println(idCountry)

	trip := models.Trip{
		Title:     request.Title,
		CountryID: request.CountryID,
		// CountryID:      idCountry.ID,
		Country: idCountry,
		// Country:        models.CountryResponse(idCountry),
		Acommodation:   request.Acommodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          dataFile,
		// UserId:         int(userId),
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// data, _ = h.TripRepository.GetTrip(data.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *tripHandlers) UpdateTrip(c echo.Context) error {
	// take Update trip response
	request := new(tripdto.UpdateTripRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	country, _ := h.TripRepository.GetCountryByID(request.CountryID)

	// modelCountry := models.CountryResponse{
	// 	ID:   0,
	// 	Name: "",
	// }

	if request.Title != "" {
		trip.Title = request.Title
	}

	trip.Country = country
	// if request.Country != modelCountry {
	// 	trip.Country = request.Country
	// }
	if request.Acommodation != "" {
		trip.Acommodation = request.Acommodation
	}
	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}
	if request.Eat != "" {
		trip.Eat = request.Eat
	}
	if request.Day != 0 {
		trip.Day = request.Day
	}
	if request.Night != 0 {
		trip.Night = request.Night
	}
	if request.DateTrip != "" {
		trip.DateTrip = request.DateTrip
	}
	if request.Price != 0 {
		trip.Price = request.Price
	}
	if request.Quota != 0 {
		trip.Quota = request.Quota
	}
	if request.Description != "" {
		trip.Description = request.Description
	}
	// if request.Image != "" {
	// 	trip.Image = request.Image
	// }
	trip.Image = path_file + trip.Image

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *tripHandlers) DeleteTrip(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TripRepository.DeleteTrip(trip, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteTrip(data)})
}

func convertResponseTrip(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		ID:    u.ID,
		Title: u.Title,
		// CountryID:      u.CountryID,
		Country:        u.Country,
		Acommodation:   u.Acommodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
		Image:          u.Image,
	}
}

func convertResponseDeleteTrip(u models.Trip) tripdto.TripDeleteResponse {
	return tripdto.TripDeleteResponse{
		ID: u.ID,
	}
}
