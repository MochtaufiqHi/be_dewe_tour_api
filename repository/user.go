package repository

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	AddUser(user models.User) (models.User, error)
	DeleteUser(ID int, user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Preload("Transaction").Find(&user).Error

	return user, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transaction").First(&user, ID).Error

	return user, err
}

func (r *repository) AddUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	// err := r.db.E
	return user, err
}

func (r *repository) DeleteUser(ID int, user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error
	// err := r.db.Delete(ID, &user).Scan(&user).Error
	// err := r.db.Exec("DELETE FROM users WHERE id=?", ID).Error
	return user, err
}
