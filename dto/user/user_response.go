package userdto

import "dumbmerch/models"

type UserResponse struct {
	ID          int                `json:"id"`
	Fullname    string             `json:"fullname"`
	Email       string             `json:"emai"`
	Password    string             `json:"password"`
	Phone       string             `json:"phone"`
	Address     string             `json:"address"`
	Transaction models.Transaction `json:"transaction"`
	Trip        models.Trip        `json:"trip"`
}

type UserDeleteResponse struct {
	ID int `json:"id"`
}
