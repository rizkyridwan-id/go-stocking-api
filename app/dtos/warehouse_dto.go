package dtos

type CreateWarehouseRequestDTO struct {
	WarehouseCode string `json:"warehouse_code" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
}

type EditWarehouseRequestDTO struct {
	WarehouseName string `json:"warehouse_name" validate:"required"`
}
