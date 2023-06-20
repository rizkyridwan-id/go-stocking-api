package repositories

import (
	"stockingapi/app/models"

	"gorm.io/gorm"
)

type ItemEventRepository struct {
	db *gorm.DB
}

func CreateItemEventRepository(db *gorm.DB) *ItemEventRepository {
	return &ItemEventRepository{db: db}
}

func (r *ItemEventRepository) Create(itemEventModel *models.ItemEventModel) *gorm.DB {
	return r.db.Create(itemEventModel)
}
