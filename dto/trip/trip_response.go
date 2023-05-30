package tripdto

import "dumbmerch/models"

type TripResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	// CountryID      int                    `json:"country_id"`
	Country        models.CountryResponse `json:"country"`
	Acommodation   string                 `json:"acommodation"`
	Transportation string                 `json:"transportation"`
	Eat            string                 `json:"eat"`
	Day            int                    `json:"day"`
	Night          int                    `json:"night"`
	DateTrip       string                 `json:"dateTrip"`
	Price          int                    `json:"price"`
	Quota          int                    `json:"quota"`
	Description    string                 `json:"description"`
	Image          string                 `json:"image"`
}

type TripDeleteResponse struct {
	ID int `json:"id"`
}
