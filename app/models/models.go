package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"autoIncrement;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Add Model Structs here, it will be auto migrate by db.go file
func LoadModels() []interface{} {
	return []interface{}{
		&UserModel{},
		&SupplierModel{},
		&WarehouseModel{},
		&ShelfModel{},
		&GroupModel{},
		&ItemModel{},
		&ItemEventModel{},
	}
}
