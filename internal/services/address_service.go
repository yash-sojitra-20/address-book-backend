package services

import (
	"errors"
	"fmt"

	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/yash-sojitra-20/address-book-backend/internal/logger"
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
	"go.uber.org/zap"
)

type IAddressService interface {
	Create(userID uint, address *models.Address) error
	GetAll(userID uint) ([]utils.AddressResponse, error)
	GetByID(userID, addressID uint) (*utils.AddressResponse, error)
	Update(userID uint, id uint, updated *models.Address) error
	Delete(userID uint, id uint) error
	ExportAddresses(userID uint, userEmail string, cfg *config.Config) error
	ExportAddressesAsync(userID uint, userEmail string, cfg *config.Config)
	ExportAddressesCustomAsync(userID uint, fields []string, sendTo string, cfg *config.Config)
	GetFilteredAddresses(userID, page, limit int, search, city, state, country string) (utils.PaginatedResponse, error)
}

type AddressService struct {
	addressRepo *repositories.AddressRepository
}

func NewAddressService(addressRepo *repositories.AddressRepository) *AddressService {
	return &AddressService{addressRepo}
}

func (s *AddressService) Create(userID uint, address *models.Address) error {
	address.UserID = userID
	return s.addressRepo.Create(address)
}

func (s *AddressService) GetAll(userID uint) ([]utils.AddressResponse, error) {
	addresses, err := s.addressRepo.FindAllByUser(userID)
	if err != nil {
		return nil, err
	}

	// Map Address to AddressResponse
	var result []utils.AddressResponse
	for _, c := range addresses {
		result = append(result, utils.AddressResponse{
			ID:           c.ID,
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			Email:        c.Email,
			Phone:        c.Phone,
			AddressLine1: c.AddressLine1,
			AddressLine2: c.AddressLine2,
			City:         c.City,
			State:        c.State,
			Country:      c.Country,
			Pincode:      c.Pincode,
		})
	}

	return result, nil
}

func (s *AddressService) GetByID(userID, addressID uint) (*utils.AddressResponse, error) {
	address, err := s.addressRepo.FindByID(addressID)
	if err != nil {
		return nil, err
	}

	if address.ID == 0 {
		return nil, errors.New("address not found")
	}

	// ownership check
	if address.UserID != userID {
		return nil, errors.New("address not found for this user")
	}

	// Map to DTO
	resp := &utils.AddressResponse{
		ID:           address.ID,
		FirstName:    address.FirstName,
		LastName:     address.LastName,
		Email:        address.Email,
		Phone:        address.Phone,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		Pincode:      address.Pincode,
	}

	return resp, nil
}

func (s *AddressService) Update(userID uint, id uint, updated *models.Address) error {
	address, err := s.addressRepo.FindByID(id)
	if err != nil {
		return err
	}

	if address.ID == 0 {
		return errors.New("address not found")
	}

	// ownership check
	if address.UserID != userID {
		return errors.New("unauthorized")
	}

	// update fields
	address.FirstName = updated.FirstName
	address.LastName = updated.LastName
	address.Email = updated.Email
	address.Phone = updated.Phone
	address.AddressLine1 = updated.AddressLine1
	address.AddressLine2 = updated.AddressLine2
	address.City = updated.City
	address.State = updated.State
	address.Country = updated.Country
	address.Pincode = updated.Pincode

	return s.addressRepo.Update(address)
}

func (s *AddressService) Delete(userID uint, id uint) error {
	address, err := s.addressRepo.FindByID(id)
	if err != nil {
		return err
	}

	if address.ID == 0 {
		return errors.New("address not found")
	}

	if address.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.addressRepo.Delete(address)
}

func (s *AddressService) ExportAddresses(
	userID uint,
	userEmail string,
	cfg *config.Config,
) error {

	// 1. Fetch addresses
	addresses, err := s.addressRepo.FindAllForExport(userID)
	if err != nil {
		return err
	}

	// 2. Generate CSV
	filePath, err := utils.GenerateAddressesCSV(userID, addresses)
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
		"Adresses CSV Export",
		"Please find attached your addresses CSV file.",
		filePath,
	)
}

func (s *AddressService) ExportAddressesAsync(
	userID uint,
	userEmail string,
	cfg *config.Config,
) {

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error(
					"panic in async export",
					zap.Any("error", r),
				)
			}
		}()

		logger.Logger.Info(
			"starting async address export",
			zap.Uint("user_id", userID),
		)

		// 1. Fetch addresses
		addresses, err := s.addressRepo.FindAllForExport(userID)
		if err != nil {
			logger.Logger.Error("failed to fetch addresses", zap.Error(err))
			return
		}

		// 2. Generate CSV
		filePath, err := utils.GenerateAddressesCSV(userID, addresses)
		if err != nil {
			logger.Logger.Error("failed to generate csv", zap.Error(err))
			return
		}

		// 3. Send email
		err = utils.SendEmailWithAttachment(
			cfg.SMTPHost,
			cfg.SMTPPort,
			cfg.SMTPUser,
			cfg.SMTPPass,
			userEmail,
			"Your addresses CSV Export",
			"Attached is your addresses CSV file.",
			filePath,
		)
		if err != nil {
			logger.Logger.Error("failed to send email", zap.Error(err))
			return
		}

		logger.Logger.Info(
			"address export completed",
			zap.Uint("user_id", userID),
			zap.String("file", filePath),
		)
	}()
}

func (s *AddressService) ExportAddressesCustomAsync(
	userID uint,
	fields []string,
	sendTo string,
	cfg *config.Config,
) {

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error("panic in async custom export", zap.Any("error", r))
			}
		}()

		logger.Logger.Info("custom address export started", zap.Uint("user_id", userID))

		// 1. Fetch all addresses for user
		addresses, err := s.addressRepo.FindAllForExport(userID)
		if err != nil {
			logger.Logger.Error("failed to fetch addresses", zap.Error(err))
			return
		}

		// 2. Convert addresses into [][]string based on requested fields
		rows := utils.FilterAddressFields(addresses, fields)

		// 3. Generate CSV from filtered data
		filePath, fileName, err := utils.GenerateCustomAddressesCSV(userID, rows)
		if err != nil {
			logger.Logger.Error("failed to generate custom csv", zap.Error(err))
			return
		}

		// Create download URL
		downloadURL := fmt.Sprintf(
			"%s/downloads/%s",
			cfg.AppURL,
			fileName,
		)

		// Email with ATTACHMENT + LINK
		emailBody := fmt.Sprintf(
			"Attached is the custom address report you requested.\n\n"+
				"You can also download it using the link below:\n%s",
			downloadURL,
		)

		// 4. Email with attachment
		err = utils.SendEmailWithAttachment(
			cfg.SMTPHost,
			cfg.SMTPPort,
			cfg.SMTPUser,
			cfg.SMTPPass,
			sendTo,
			"Custom Address CSV Export",
			emailBody,
			filePath,
		)
		if err != nil {
			logger.Logger.Error("failed to send export email", zap.Error(err))
			return
		}

		logger.Logger.Info(
			"custom address export completed",
			zap.Uint("user_id", userID),
			zap.String("file", filePath),
		)
	}()
}

func (s *AddressService) GetFilteredAddresses(
	userID, page, limit int,
	search, city, state, country string,
) (utils.PaginatedResponse, error) {

	addresses, total, err := s.addressRepo.FindFiltered(
		uint(userID),
		page, limit,
		search, city, state, country,
	)

	if err != nil {
		// fmt.Println("inside service:",err)
		return utils.PaginatedResponse{}, err
	}

	var responseData []utils.AddressResponse
	for _, a := range addresses {
		responseData = append(responseData, utils.AddressResponse{
			ID:           a.ID,
			FirstName:    a.FirstName,
			LastName:     a.LastName,
			Email:        a.Email,
			Phone:        a.Phone,
			AddressLine1: a.AddressLine1,
			AddressLine2: a.AddressLine2,
			City:         a.City,
			State:        a.State,
			Country:      a.Country,
			Pincode:      a.Pincode,
		})
	}

	return utils.PaginatedResponse{
		Page:  page,
		Limit: limit,
		Total: total,
		Data:  responseData,
	}, nil
}
