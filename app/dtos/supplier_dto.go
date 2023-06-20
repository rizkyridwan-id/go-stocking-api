package dtos

type CreateSupplierRequestDTO struct {
	SupplierCode string `json:"supplier_code" validate:"required"`
	SupplierName string `json:"supplier_name" validate:"required"`
}

type EditSupplierRequestDTO struct {
	SupplierName string `json:"supplier_name" validate:"required"`
}
