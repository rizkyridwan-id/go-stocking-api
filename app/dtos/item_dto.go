package dtos

import (
	"stockingapi/app/helpers"
)

type CreateItemRequestDTO struct {
	GroupCode    string  `json:"group_code" validate:"required"`
	ShelfCode    string  `json:"shelf_code" validate:"required"`
	SupplierCode string  `json:"supplier_code" validate:"required"`
	ItemCode     string  `json:"item_code" validate:"required"`
	ItemName     string  `json:"item_name" validate:"required"`
	ItemInfo     string  `json:"item_info"`
	Size         string  `json:"size"`
	Stock        uint    `json:"stock" validate:"required"`
	Weight       float64 `json:"weight" validate:"required"`
	Price        uint    `json:"price" validate:"required"`
}

type TransactionItemRequestDTO struct {
	ItemCode string `json:"item_code" validate:"required"`
	Stock    uint   `json:"stock" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

func (d *TransactionItemRequestDTO) ValidateType() bool {
	validType := []string{"IN", "OUT"}
	return helpers.ContainsPrimitive(validType, d.Type)
}

type MovementItemRequestDTO struct {
	ItemCode    string `json:"item_code" validate:"required"`
	ShelfBefore string `json:"shelf_code_before" validate:"required"`
	ShelfAfter  string `json:"shelf_code_after" validate:"required"`
}
