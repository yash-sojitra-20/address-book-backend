package repositories

import (
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db}
}

func (r *ContactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *ContactRepository) FindAllByUser(userID uint) ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Where("user_id = ?", userID).Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) FindByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.First(&contact, id).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *ContactRepository) Update(contact *models.Contact) error {
	return r.db.Save(contact).Error
}

func (r *ContactRepository) Delete(contact *models.Contact) error {
	return r.db.Delete(contact).Error // soft delete
}

// Query format:
// GET /contacts?page=1&limit=10&city=Ahmedabad
func (r *ContactRepository) FindPaginated(
	userID uint,
	page int,
	limit int,
	city string,
) ([]models.Contact, error) {

	offset := (page - 1) * limit

	query := r.db.Where("user_id = ?", userID)

	if city != "" {
		query = query.Where("city = ?", city)
	}

	var contacts []models.Contact
	err := query.
		Limit(limit).
		Offset(offset).
		Find(&contacts).Error

	return contacts, err
}
