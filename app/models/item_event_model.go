package models

import "time"

type ItemEventModel struct {
	BaseModel
	RItemCode   string
	RShelfCode  string
	Type        string     `gorm:"not null"`
	StockBefore uint       `gorm:"not null"`
	StockChange int        `gorm:"not null"`
	StockAfter  uint       `gorm:"not null"`
	CreatedDate time.Time  `gorm:"type:date;default:CURRENT_DATE"`
	CreatedBy   string     `gorm:"not null"`
	Item        ItemModel  `gorm:"foreignKey:RItemCode;references:ItemCode"`
	Shelf       ShelfModel `gorm:"foreignKey:RShelfCode;references:ShelfCode"`
}

func (ItemEventModel) TableName() string {
	return "tt_item_event"
}
