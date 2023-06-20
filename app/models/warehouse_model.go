package models

type WarehouseModel struct {
	BaseModel
	WarehouseCode string `gorm:"primaryKey"`
	WarehouseName string `gorm:"not null"`
}

func (WarehouseModel) TableName() string {
	return "tm_warehouse"
}

type EditWarehouseModel struct {
	WarehouseName string
}
