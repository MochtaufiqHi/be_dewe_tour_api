package trandto

import "dumbmerch/models"

type CreateTransactionRequest struct {
	CounterQty int                 `json:"counterQty" validate:"required"`
	Total      int                 `json:"total" validate:"required"`
	Status     string              `json:"status" validate:"required"`
	Attachment string              `json:"attachment"`
	TripID     int                 `json:"trip_id" `
	Trip       models.TripResponse `json:"trip"`
	UserID     int                 `json:"user_id"`
	User       models.User         `json:"user"`
}
