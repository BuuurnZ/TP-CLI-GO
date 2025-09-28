package stores

import (
	"encoding/json"
	"errors"
	"loganalyzer/internal/models"
	"os"
	"sync"
)

type JSONStore struct {
	filePath string
	contacts []models.Contact
	mu       sync.RWMutex
}

func NewJSONStore(filePath string) (*JSONStore, error) {
	store := &JSONStore{
		filePath: filePath,
		contacts: make([]models.Contact, 0),
	}

	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *JSONStore) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		s.contacts = make([]models.Contact, 0)
		return nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		s.contacts = make([]models.Contact, 0)
		return nil
	}

	return json.Unmarshal(data, &s.contacts)
}

func (s *JSONStore) save() error {
	data, err := json.MarshalIndent(s.contacts, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(s.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *JSONStore) Create(contact *models.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	maxID := uint(0)
	for _, c := range s.contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	contact.ID = maxID + 1
	s.contacts = append(s.contacts, *contact)
	return s.save()
}

func (s *JSONStore) GetByID(id uint) (*models.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, contact := range s.contacts {
		if contact.ID == id {
			return &contact, nil
		}
	}
	return nil, errors.New("record not found")
}

func (s *JSONStore) GetAll() ([]models.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.contacts, nil
}

func (s *JSONStore) Update(contact *models.Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, c := range s.contacts {
		if c.ID == contact.ID {
			s.contacts[i] = *contact
			return s.save()
		}
	}
	return errors.New("record not found")
}

func (s *JSONStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, contact := range s.contacts {
		if contact.ID == id {
			s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
			return s.save()
		}
	}
	return errors.New("record not found")
}

func (s *JSONStore) GetByEmail(email string) (*models.Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, contact := range s.contacts {
		if contact.Email == email {
			return &contact, nil
		}
	}
	return nil, errors.New("record not found")
}
