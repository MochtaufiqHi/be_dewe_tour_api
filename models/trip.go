package models

type Trip struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Title     string `json:"title" form:"title" gorm:"type: varchar(255)"`
	CountryID int    `json:"country_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Country        CountryResponse `json:"country" gorm:"foreignKey:CountryID"`
	Country        CountryResponse `json:"country" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Acommodation   string          `json:"acommodation" form:"acommodation" gorm:"type: varchar(255)"`
	Transportation string          `json:"transportation" form:"transportation" gorm:"type: varchar(255)"`
	Eat            string          `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Day            int             `json:"day" form:"day" gorm:"type: int"`
	Night          int             `json:"night" form:"night" gorm:"type: int"`
	DateTrip       string          `json:"dateTrip" form:"dateTrip" gorm:"type: varchar(255)"`
	Price          int             `json:"price" form:"price" gorm:"type: int"`
	Quota          int             `json:"quota" form:"quota" gorm:"type: int"`
	Description    string          `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image          string          `json:"image" form:"image" gorm:"type: varchar(255)"`
	// UserId         int             `json:"user_id" from:"user_id"`
	// User           UserResponse    `json:"user"`
}

type TripResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	CountryID      int             `json:"country_id" gorm:"constrain:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Country        CountryResponse `json:"country" gorm:"foreignKey:CountryID"`
	Acommodation   string          `json:"acommodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat" form:"eat"`
	Day            int             `json:"day" form:"day"`
	Night          int             `json:"night" form:"night"`
	DateTrip       string          `json:"dateTrip"`
	Price          int             `json:"price" form:"price"`
	Quota          int             `json:"quota" form:"quota"`
	Description    string          `json:"description"`
	Image          string          `json:"image"`
}

func (TripResponse) TableName() string {
	return "trips"
}
