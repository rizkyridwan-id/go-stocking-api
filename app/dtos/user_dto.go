package dtos

type RegisterRequestDTO struct {
	UserName string `json:"user_name" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	Level    string `json:"level" validate:"required"`
}

type EditRequestDTO struct {
	UserName string `json:"user_name" validate:"required"`
}

type LoginRequestDTO struct {
	UserId   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	UserId  string `json:"user_id"`
	Level   string `json:"level"`
}
