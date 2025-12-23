package services

import (
	"errors"

	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
)

type ContactService struct {
	contactRepo *repositories.ContactRepository
}

func NewContactService(contactRepo *repositories.ContactRepository) *ContactService {
	return &ContactService{contactRepo}
}

func (s *ContactService) Create(userID uint, contact *models.Contact) error {
	contact.UserID = userID
	return s.contactRepo.Create(contact)
}

func (s *ContactService) GetAll(userID uint) ([]models.Contact, error) {
	return s.contactRepo.FindAllByUser(userID)
}


func (s *ContactService) Update(userID uint, id uint, updated *models.Contact) error {
	contact, err := s.contactRepo.FindByID(id)
	if err != nil {
		return errors.New("contact not found")
	}

	// ownership check
	if contact.UserID != userID {
		return errors.New("unauthorized")
	}

	// update fields
	contact.FirstName = updated.FirstName
	contact.LastName = updated.LastName
	contact.Email = updated.Email
	contact.Phone = updated.Phone
	contact.AddressLine1 = updated.AddressLine1
	contact.AddressLine2 = updated.AddressLine2
	contact.City = updated.City
	contact.State = updated.State
	contact.Country = updated.Country
	contact.Pincode = updated.Pincode

	return s.contactRepo.Update(contact)
}

func (s *ContactService) Delete(userID uint, id uint) error {
	contact, err := s.contactRepo.FindByID(id)
	if err != nil {
		return errors.New("contact not found")
	}

	if contact.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.contactRepo.Delete(contact)
}