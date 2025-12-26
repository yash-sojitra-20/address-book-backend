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

func (r *AddressRepository) FindAllForExport(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.
		Where("user_id = ?", userID).
		Find(&addresses).Error
	return addresses, err
}

// Query format:
// GET /addreses?page=1&limit=10&city=Ahmedabad
func (r *AddressRepository) FindFiltered(
	userID uint,
	page int,
	limit int,
	search string,
	city string,
	state string,
	country string,
) ([]models.Address, int64, error) {

	offset := (page - 1) * limit

	query := r.db.Model(&models.Address{}).Where("user_id = ?", userID)

	// SEARCH (across multiple fields)
	if search != "" {
		like := "%" + search + "%"
		query = query.Where(`
			first_name ILIKE ? OR 
			last_name ILIKE ? OR 
			email ILIKE ? OR
			phone ILIKE ? OR
			city ILIKE ? OR
			state ILIKE ? OR
			country ILIKE ?`,
			like, like, like, like, like, like, like,
		)
	}

	// FILTERS
	if city != "" {
		query = query.Where("city ILIKE ?", city)
	}
	if state != "" {
		query = query.Where("state ILIKE ?", state)
	}
	if country != "" {
		query = query.Where("country ILIKE ?", country)
	}

	// fmt.Println(query)

	var total int64
	query.Count(&total) // get total records

	// fmt.Println(total)

	// PAGINATION
	var addresses []models.Address
	err := query.
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&addresses).Error

	// fmt.Println("inside repo:",err)

	return addresses, total, err
}
