package models

import "time"

type ItemModel struct {
	BaseModel
	RGroupCode    string        `json:"r_group_code"`
	RShelfCode    string        `json:"r_shelf_code"`
	RSupplierCode string        `json:"r_supplier_code"`
	ItemCode      string        `json:"item_code" gorm:"primaryKey"`
	ItemName      string        `json:"item_name" gorm:"not null"`
	ItemInfo      string        `json:"item_info"`
	Size          string        `json:"size"`
	Stock         uint          `json:"stock" gorm:"not null"`
	Weight        float64       `json:"weight" gorm:"not null"`
	Price         uint          `json:"price" gorm:"not null"`
	CreatedDate   time.Time     `json:"created_date" gorm:"type:date;default:CURRENT_DATE"`
	CreatedBy     string        `json:"created_by" gorm:"not null"`
	Group         GroupModel    `gorm:"foreignKey:RGroupCode;references:GroupCode"`
	Shelf         ShelfModel    `gorm:"foreignKey:RShelfCode;references:ShelfCode"`
	Supplier      SupplierModel `gorm:"foreignKey:RSupplierCode;references:SupplierCode"`
}

func (ItemModel) TableName() string {
	return "tm_item"
}
