package repositories

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	FindTicket() ([]models.Ticket, error)
	GetTicket(id int) (models.Ticket, error)
}

func RepositoryTicket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) FindTicket() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").Find(&tickets).Error

	return tickets, err
}

func (r repository) GetTicket(ID int) (models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Preload("StartStation").Preload("DestinationStation").First(&ticket, ID).Error

	return ticket, err
}
