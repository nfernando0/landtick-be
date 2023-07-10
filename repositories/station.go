package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type StationRepository interface {
	FindStation() ([]models.Station, error)
	GetStation(id int) (models.Station, error)
	CreateStation(station models.Station) (models.Station, error)
	UpdateStation(station models.Station) (models.Station, error)
	DeleteStation(station models.Station) (models.Station, error)
}

type repositoryStation struct {
	db *gorm.DB
}

func RepositoryStation(db *gorm.DB) *repositoryStation {
	return &repositoryStation{db}
}

func (r *repositoryStation) FindStation() ([]models.Station, error) {
	var stations []models.Station
	err := r.db.Find(&stations).Error

	return stations, err
}

func (r *repositoryStation) GetStation(id int) (models.Station, error) {
	var station models.Station
	err := r.db.First(&station, id).Error

	return station, err
}

func (r *repositoryStation) CreateStation(station models.Station) (models.Station, error) {
	err := r.db.Create(&station).Error

	return station, err
}
func (r *repositoryStation) UpdateStation(station models.Station) (models.Station, error) {
	err := r.db.Save(&station).Error

	return station, err
}

func (r *repositoryStation) DeleteStation(station models.Station) (models.Station, error) {
	err := r.db.Delete(&station).Scan(&station).Error

	return station, err
}
