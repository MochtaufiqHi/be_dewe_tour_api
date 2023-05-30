package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	GetAllTrip() ([]models.Trip, error)
	GetTrip(ID int) (models.Trip, error)
	GetCountryByID(ID int) (models.CountryResponse, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	UpdateTrip(trip models.Trip) (models.Trip, error)
	DeleteTrip(trip models.Trip, ID int) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTrip() ([]models.Trip, error) {
	var trip []models.Trip
	err := r.db.Preload("Country").Find(&trip).Error

	return trip, err
}

func (r *repository) GetTrip(ID int) (models.Trip, error) {
	var trip models.Trip
	err := r.db.Preload("Country").First(&trip, ID).Error

	return trip, err
}

func (r *repository) UpdateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Save(&trip).Error

	return trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Create(&trip).Error

	return trip, err
}

func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {
	err := r.db.Delete(&trip).Error

	return trip, err
}

func (r *repository) GetCountryByID(ID int) (models.CountryResponse, error) {
	var country models.CountryResponse

	err := r.db.First(&country, ID).Error

	return country, err
}
