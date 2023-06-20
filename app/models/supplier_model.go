package models

type SupplierModel struct {
	BaseModel
	SupplierCode string `gorm:"primaryKey"`
	SupplierName string `gorm:"not null"`
}

func (SupplierModel) TableName() string {
	return "tm_supplier"
}

type EditSupplierModel struct {
	SupplierName string
}
