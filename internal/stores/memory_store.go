package stores

import (
	"errors"
	"loganalyzer/internal/models"
	"sync"
)

type MemoryStore struct {
	contacts []models.Contact
	nextID   uint
	mu       sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make([]models.Contact, 0),
		nextID:   1,
	}
}

func (s *MemoryStore) Create(contact *models.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact.ID = s.nextID
	s.nextID++
	s.contacts = append(s.contacts, *contact)
	return nil
}

func (s *MemoryStore) GetByID(id uint) (*models.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, contact := range s.contacts {
		if contact.ID == id {
			return &contact, nil
		}
	}
	return nil, errors.New("record not found")
}

func (s *MemoryStore) GetAll() ([]models.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.contacts, nil
}

func (s *MemoryStore) Update(contact *models.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, c := range s.contacts {
		if c.ID == contact.ID {
			s.contacts[i] = *contact
			return nil
		}
	}
	return errors.New("record not found")
}

func (s *MemoryStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, contact := range s.contacts {
		if contact.ID == id {
			s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
			return nil
		}
	}
	return errors.New("record not found")
}

func (s *MemoryStore) GetByEmail(email string) (*models.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, contact := range s.contacts {
		if contact.Email == email {
			return &contact, nil
		}
	}
	return nil, errors.New("record not found")
}
