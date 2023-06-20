package models

type ShelfModel struct {
	BaseModel
	ShelfCode      string         `json:"shelf_code" gorm:"primaryKey"`
	ShelfName      string         `json:"shelf_name" gorm:"not null"`
	RWarehouseCode string         `json:"r_warehouse_code"`
	Warehouse      WarehouseModel `gorm:"foreignKey:RWarehouseCode;references:WarehouseCode"`
}

func (ShelfModel) TableName() string {
	return "tm_shelf"
}

type EditShelfModel struct {
	ShelfName string
}
