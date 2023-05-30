package trandto

import "dumbmerch/models"

type TransactionResponse struct {
	ID         int         `json:"id"`
	CounterQty int         `json:"counterQty"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	Attachment string      `json:"attachment"`
	TripID     int         `json:"trip_id"`
	Trip       models.Trip `json:"trip"`
	User       models.User `json:"user"`
}

type TransactionUpdateResponse struct {
	CounterQty int         `json:"counterQty"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	Attachment string      `json:"attachment"`
	TripID     int         `json:"trip_id"`
	Trip       models.Trip `json:"trip"`
}
