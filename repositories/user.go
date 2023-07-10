package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUser() ([]models.UserResponse, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser() ([]models.UserResponse, error) {
	var users []models.UserResponse
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user, ID).Error
	return user, err
}
