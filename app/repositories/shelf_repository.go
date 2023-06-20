package repositories

import (
	"errors"
	"fmt"
	"stockingapi/app/models"

	"gorm.io/gorm"
)

type ShelfRepository struct {
	db *gorm.DB
}

func CreateShelfRepository(db *gorm.DB) *ShelfRepository {
	return &ShelfRepository{
		db: db,
	}
}

func (r *ShelfRepository) FindAll() (*[]models.ShelfModel, error) {
	var shelfModels []models.ShelfModel

	result := r.db.Preload("Warehouse").Find(&shelfModels)
	if result.Error != nil {
		fmt.Println("ERR: (SupplierRepository)(FindAll) ", result.Error.Error())
		return nil, result.Error
	}

	return &shelfModels, nil
}

func (r *ShelfRepository) Create(shelfModel *models.ShelfModel) *gorm.DB {
	return r.db.Create(&shelfModel)
}

func (r *ShelfRepository) FindOneShelfByCode(shelfCode string) (*models.ShelfModel, error) {
	var shelfModel models.ShelfModel

	result := r.db.Where("shelf_code = ?", shelfCode).Preload("Warehouse").First(&shelfModel)
	if result.Error != nil {
		fmt.Println("ERR: (ShelfRepository)(FindOneShelfByCode) ", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &shelfModel, nil
}

func (r *ShelfRepository) UpdateByCode(shelfCode string, updateModel *models.EditShelfModel) error {
	var shelfModel models.ShelfModel

	result := r.db.Where("shelf_code = ?", shelfCode).First(&shelfModel)
	if result.Error != nil {
		return result.Error
	}

	shelfModel.ShelfName = updateModel.ShelfName

	r.db.Save(&shelfModel)
	return nil
}

func (r *ShelfRepository) DeleteByCode(shelfCode string) error {
	result := r.db.Where("shelf_code = ?", shelfCode).Delete(&models.ShelfModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
