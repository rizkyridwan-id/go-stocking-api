package repositories

import (
	"errors"
	"fmt"
	"stockingapi/app/models"
	"time"

	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func CreateItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) Create(itemModel *models.ItemModel) *gorm.DB {
	return r.db.Create(itemModel)
}

func (r *ItemRepository) FindBy() (*[]models.ItemModel, error) {
	var itemModels []models.ItemModel

	result := r.db.Preload("Shelf").Preload("Shelf.Warehouse").Preload("Group").Preload("Supplier").Find(&itemModels)
	if result.Error != nil {
		return nil, result.Error
	}

	return &itemModels, nil
}

func (r *ItemRepository) FindOneByCode(itemCode string) (*models.ItemModel, error) {
	var ItemModel models.ItemModel

	result := r.db.Where("item_code = ?", itemCode).Preload("Shelf").First(&ItemModel)
	if result.Error != nil {
		fmt.Println("ERR: (ItemRepository)(FindOneByCode) ", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &ItemModel, nil
}

func (r *ItemRepository) UpdateStock(itemCode string, stock uint) error {
	var itemModel models.ItemModel

	result := r.db.Where("item_code = ?", itemCode).First(&itemModel)
	if result.Error != nil {
		return result.Error
	}

	itemModel.Stock = stock
	itemModel.UpdatedAt = time.Now()

	r.db.Save(&itemModel)
	return nil
}

func (r *ItemRepository) UpdateShelf(itemCode string, shelfCode string) error {
	var itemModel models.ItemModel

	result := r.db.Where("item_code = ?", itemCode).First(&itemModel)
	if result.Error != nil {
		return result.Error
	}

	itemModel.RShelfCode = shelfCode
	itemModel.UpdatedAt = time.Now()

	r.db.Save(&itemModel)
	return nil
}
