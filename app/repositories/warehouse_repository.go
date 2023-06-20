package repositories

import (
	"errors"
	"fmt"
	"stockingapi/app/models"
	"time"

	"gorm.io/gorm"
)

type WarehouseRepository struct {
	db *gorm.DB
}

func CreateWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{
		db: db,
	}
}

func (r *WarehouseRepository) FindAll() (*[]models.WarehouseModel, error) {
	var warehouseModels []models.WarehouseModel

	result := r.db.Find(&warehouseModels)
	if result.Error != nil {
		fmt.Println("ERR: (WarehouseRepository)(FindAll) ", result.Error.Error())
		return nil, result.Error
	}

	return &warehouseModels, nil
}

func (r *WarehouseRepository) Create(warehouseModel *models.WarehouseModel) *gorm.DB {
	return r.db.Create(warehouseModel)
}

func (r *WarehouseRepository) FindOneByWarehouseCode(warehouseCode string) (*models.WarehouseModel, error) {
	var warehouseModel models.WarehouseModel

	result := r.db.Where("warehouse_code = ?", warehouseCode).First(&warehouseModel)
	if result.Error != nil {
		fmt.Printf("ERR: (WarehouseRepository)(FindOneByWarehouseCode) %s", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &warehouseModel, nil
}

func (r *WarehouseRepository) UpdateByCode(warehouseCode string, updateModel *models.EditWarehouseModel) error {
	var warehouseModel models.WarehouseModel

	result := r.db.Where("warehouse_code = ?", warehouseCode).First(&warehouseModel)
	if result.Error != nil {
		return result.Error
	}

	warehouseModel.WarehouseName = updateModel.WarehouseName
	warehouseModel.UpdatedAt = time.Now()

	r.db.Save(&warehouseModel)
	return nil
}

func (r *WarehouseRepository) DeleteByCode(supplierCode string) error {
	result := r.db.Where("warehouse_code = ?", supplierCode).Delete(&models.WarehouseModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
