package models

type Transaction struct {
	ID         int          `json:"id"`
	CounterQty int          `json:"counterQty" gorm:"type: int"`
	Total      int          `json:"total" gorm:"type: int"`
	Status     string       `json:"status" gorm:"type: varchar(255)"`
	Attachment string       `json:"attachment" gorm:"type: varchar(255)"`
	UserID     int          `json:"user_id" gorm:"type: int"`
	User       User         `json:"user" gorm:"foreignKey:UserID"`
	TripID     int          `json:"trip_id" gorm:"type: int"`
	Trip       TripResponse `json:"trip" gorm:"foreignKey:TripID"`
}

type TransactionResponse struct {
	ID         int    `json:"id"`
	CounterQty int    `json:"counterQty" gorm:"type: int"`
	Total      int    `json:"total" gorm:"type: int"`
	Status     string `json:"status" gorm:"type: varchar(255)"`
	Attachment string `json:"attachment" gorm:"type: varchar(255)"`
	UserID     int    `json:"user_id"`
	TripID     int    `json:"trip_id" gorm:"type: int"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
