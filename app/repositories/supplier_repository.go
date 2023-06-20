package repositories

import (
	"errors"
	"fmt"
	"stockingapi/app/models"
	"time"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func CreateSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) FindAll() (*[]models.SupplierModel, error) {
	var supplierModels []models.SupplierModel

	result := r.db.Find(&supplierModels)
	if result.Error != nil {
		fmt.Println("ERR: (SupplierRepository)(FindAll) ", result.Error.Error())
		return nil, result.Error
	}

	return &supplierModels, nil
}

func (r *SupplierRepository) Create(supplierModel *models.SupplierModel) *gorm.DB {
	return r.db.Create(supplierModel)
}

func (r *SupplierRepository) FindOneSupplierByCode(supplierCode string) (*models.SupplierModel, error) {
	var supplierModel models.SupplierModel

	result := r.db.Where("supplier_code = ?", supplierCode).First(&supplierModel)
	if result.Error != nil {
		fmt.Println("ERR: (supplierRepository)(FindOneSupplierByCode)", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &supplierModel, nil
}

func (r *SupplierRepository) UpdateByCode(supplierCode string, updateModel *models.EditSupplierModel) error {
	var supplierModel models.SupplierModel

	result := r.db.Where("supplier_code = ?", supplierCode).First(&supplierModel)
	if result.Error != nil {
		return result.Error
	}

	supplierModel.SupplierName = updateModel.SupplierName
	supplierModel.UpdatedAt = time.Now()

	r.db.Save(&supplierModel)
	return nil
}

func (r *SupplierRepository) DeleteByCode(supplierCode string) error {
	result := r.db.Where("supplier_code = ?", supplierCode).Delete(&models.SupplierModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
