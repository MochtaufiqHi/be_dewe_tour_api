package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	GetAllCountry() ([]models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	GetCountry(ID int) (models.Country, error)
	UpdateCountry(country models.Country) (models.Country, error)
	DeleteCountry(country models.Country, ID int) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllCountry() ([]models.Country, error) {
	var country []models.Country
	err := r.db.Find(&country).Error
	return country, err
}

func (r *repository) GetCountry(ID int) (models.Country, error) {
	var country models.Country

	err := r.db.First(&country, ID).Error

	return country, err
}

func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error

	return country, err
}

func (r *repository) UpdateCountry(country models.Country) (models.Country, error) {
	// err := r.db.Exec("UPDATE countries WHERE id=?", ID).Error
	err := r.db.Save(&country).Error

	return country, err
}

func (r *repository) DeleteCountry(country models.Country, ID int) (models.Country, error) {
	// err := r.db.Exec("DELETE FROM countries WHERE id=?", ID).Scan(&country).Error
	err := r.db.Delete(&country).Error

	return country, err
}
