package stores

import "loganalyzer/internal/models"

type Storer interface {
	Create(contact *models.Contact) error
	GetByID(id uint) (*models.Contact, error)
	GetAll() ([]models.Contact, error)
	Update(contact *models.Contact) error
	Delete(id uint) error
	GetByEmail(email string) (*models.Contact, error)
}
