package tripdto

import "dumbmerch/models"

type CreateTripRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title" form:"title"`
	CountryID int    `json:"country_id" form:"country_"`
	// Country        models.CountryResponse `json:"country" validate:"required"`
	Acommodation   string `json:"acommodation" form:"acommodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Eat            string `json:"eat" form:"eat"`
	Day            int    `json:"day" form:"day"`
	Night          int    `json:"night" form:"night"`
	DateTrip       string `json:"dateTrip" form:"dateTrip"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	Description    string `json:"description" form:"description"`
	Image          string `json:"image" form:"image"`
}

type UpdateTripRequest struct {
	ID             int                    `json:"id"`
	Title          string                 `json:"title" form:"title" validate:"required"`
	CountryID      int                    `json:"country_id" form:"country_id" validate:"required"`
	Country        models.CountryResponse `json:"country" validate:"required"`
	Acommodation   string                 `json:"acommodation" form:"acommodation" validate:"required"`
	Transportation string                 `json:"transportation" form:"transportation" validate:"required"`
	Eat            string                 `json:"eat" form:"eat" validate:"required"`
	Day            int                    `json:"day" form:"day" validate:"required"`
	Night          int                    `json:"night" form:"night" validate:"required"`
	DateTrip       string                 `json:"dateTrip" form:"dateTrip" validate:"required"`
	Price          int                    `json:"price" form:"price" validate:"required"`
	Quota          int                    `json:"quota" form:"quota" validate:"required"`
	Description    string                 `json:"description" form:"description" validate:"required"`
	Image          string                 `json:"image" form:"image" validate:"required"`
}
