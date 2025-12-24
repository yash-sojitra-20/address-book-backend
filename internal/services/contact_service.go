package services

import (
	"errors"

	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/yash-sojitra-20/address-book-backend/internal/middleware"
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
	"go.uber.org/zap"
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

func (s *ContactService) GetPaginated(
	userID uint,
	page int,
	limit int,
	city string,
) ([]models.Contact, error) {
	return s.contactRepo.FindPaginated(userID, page, limit, city)
}

func (s *ContactService) ExportContacts(
	userID uint,
	userEmail string,
	cfg *config.Config,
) error {

	// 1. Fetch contacts
	contacts, err := s.contactRepo.FindAllForExport(userID)
	if err != nil {
		return err
	}

	// 2. Generate CSV
	filePath, err := utils.GenerateContactsCSV(userID, contacts)
	if err != nil {
		return err
	}

	// 3. Send email with attachment
	return utils.SendEmailWithAttachment(
		cfg.SMTPHost,
		cfg.SMTPPort,
		cfg.SMTPUser,
		cfg.SMTPPass,
		userEmail,
		"Contacts CSV Export",
		"Please find attached your contacts CSV file.",
		filePath,
	)
}

func (s *ContactService) ExportContactsAsync(
	userID uint,
	userEmail string,
	cfg *config.Config,
) {

	go func() {
		defer func() {
			if r := recover(); r != nil {
				middleware.Logger.Error(
					"panic in async export",
					zap.Any("error", r),
				)
			}
		}()

		middleware.Logger.Info(
			"starting async contact export",
			zap.Uint("user_id", userID),
		)

		// 1. Fetch contacts
		contacts, err := s.contactRepo.FindAllForExport(userID)
		if err != nil {
			middleware.Logger.Error("failed to fetch contacts", zap.Error(err))
			return
		}

		// 2. Generate CSV
		filePath, err := utils.GenerateContactsCSV(userID, contacts)
		if err != nil {
			middleware.Logger.Error("failed to generate csv", zap.Error(err))
			return
		}

		// 3. Send email
		err = utils.SendEmailWithAttachment(
			cfg.SMTPHost,
			cfg.SMTPPort,
			cfg.SMTPUser,
			cfg.SMTPPass,
			userEmail,
			"Your Contacts CSV Export",
			"Attached is your contacts CSV file.",
			filePath,
		)
		if err != nil {
			middleware.Logger.Error("failed to send email", zap.Error(err))
			return
		}

		middleware.Logger.Info(
			"contact export completed",
			zap.Uint("user_id", userID),
			zap.String("file", filePath),
		)
	}()
}
