package repositories

import (
	"errors"

	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email=?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	// var user models.User
	// result := r.db.Where("email = ?", email).Find(&user)

	// if result.Error != nil {
	// 	// real DB error
	// 	return false, result.Error
	// }

	// // check if any row matched
	// return result.RowsAffected > 0, nil

	var count int64
	err := r.db.Model(&models.User{}).Where("email=?", email).Count(&count).Error
	return count > 0, err
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
