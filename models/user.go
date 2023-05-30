package models

import "time"

type User struct {
	ID          int                   `json:"id"`
	Fullname    string                `json:"fullname" gorm:"type: varchar(255)"`
	Email       string                `json:"email" gorm:"type: varchar(255)"`
	Password    string                `json:"password" gorm:"type: varchar(255)"`
	Phone       string                `json:"phone" gorm:"type: varchar(255)"`
	Address     string                `json:"address" gorm:"type: text"`
	Role        string                `json:"role" gorm:"type: varchar(255)"`
	Transaction []TransactionResponse `json:"transaction"`
	CreatedAt   time.Time             `json:"-"`
	UpdatedAt   time.Time             `json:"-"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: text"`
}

func (UserResponse) TableName() string {
	return "users"
}
