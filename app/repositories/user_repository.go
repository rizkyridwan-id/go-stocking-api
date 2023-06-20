package repositories

import (
	"errors"
	"fmt"
	"stockingapi/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindBy() ([]*models.UserModel, error) {
	var user []*models.UserModel
	result := r.db.Where("is_deleted = ?", false).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) FindOneByUserId(userId string) (*models.UserModel, error) {
	var userModel *models.UserModel
	result := r.db.Where("user_id = ?", userId).First(&userModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		fmt.Printf("ERR: (UserRepository)(FindOneByUserId) %s\n", result.Error.Error())
		return nil, result.Error
	}

	return userModel, nil
}

func (r *UserRepository) Create(userModel *models.UserModel) *gorm.DB {
	return r.db.Create(userModel)
}
