package dtos

type CreateGroupRequestDTO struct {
	GroupCode string `json:"group_code" validate:"required"`
	GroupName string `json:"group_name" validate:"required"`
}

type EditGroupRequestDTO struct {
	GroupName string `json:"group_name" validate:"required"`
}
