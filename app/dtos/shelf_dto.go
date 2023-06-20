package dtos

type CreateShelfRequestDTO struct {
	ShelfCode     string `json:"shelf_code" validate:"required"`
	ShelfName     string `json:"shelf_name" validate:"required"`
	WarehouseCode string `json:"warehouse_code" validate:"required"`
}

type EditShelfRequestDTO struct {
	ShelfName string `json:"shelf_name" validate:"required"`
}
