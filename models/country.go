package models

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
