package stores

import (
	"loganalyzer/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Contact{}); err != nil {
		return nil, err
	}

	return &GORMStore{db: db}, nil
}

func (s *GORMStore) Create(contact *models.Contact) error {
	return s.db.Create(contact).Error
}

func (s *GORMStore) GetByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := s.db.First(&contact, id).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (s *GORMStore) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := s.db.Find(&contacts).Error
	return contacts, err
}

func (s *GORMStore) Update(contact *models.Contact) error {
	return s.db.Save(contact).Error
}

func (s *GORMStore) Delete(id uint) error {
	return s.db.Delete(&models.Contact{}, id).Error
}

func (s *GORMStore) GetByEmail(email string) (*models.Contact, error) {
	var contact models.Contact
	err := s.db.Where("email = ?", email).First(&contact).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}
