package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransaction() ([]models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	GetTripByID(ID int) (models.TripResponse, error)
	GetUserByID(ID int) (models.User, error)
	GetTransactionByUser(ID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("User").Preload("Trip.Country").Find(&transaction).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip.Country").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("user_id =?", ID).Preload("User").Preload("Trip.Country").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Trip.Country").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip.Country").Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}

func (r *repository) GetTripByID(ID int) (models.TripResponse, error) {
	var trip models.TripResponse
	err := r.db.Preload("Country").First(&trip, ID).Error

	return trip, err
}

func (r *repository) GetUserByID(ID int) (models.User, error) {
	var trip models.User
	err := r.db.Preload("User").First(&trip, ID).Error

	return trip, err
}
