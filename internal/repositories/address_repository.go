package repositories

import (
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db}
}

func (r *AddressRepository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *AddressRepository) FindAllByUser(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *AddressRepository) FindByID(id uint) (*models.Address, error) {
	var address models.Address
	err := r.db.First(&address, id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *AddressRepository) Update(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *AddressRepository) Delete(address *models.Address) error {
	return r.db.Delete(address).Error // soft delete
}

// Query format:
// GET /addreses?page=1&limit=10&city=Ahmedabad
func (r *AddressRepository) FindPaginated(
	userID uint,
	page int,
	limit int,
	city string,
) ([]models.Address, error) {

	offset := (page - 1) * limit

	query := r.db.Where("user_id = ?", userID)

	if city != "" {
		query = query.Where("city = ?", city)
	}

	var addresses []models.Address
	err := query.
		Limit(limit).
		Offset(offset).
		Find(&addresses).Error

	return addresses, err
}

func (r *AddressRepository) FindAllForExport(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.
		Where("user_id = ?", userID).
		Find(&addresses).Error
	return addresses, err
}
