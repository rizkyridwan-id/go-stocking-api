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
	Message               string `json:"message"`
	UserId                string `json:"user_id"`
	Level                 string `json:"level"`
	Token                 string `json:"token"`
	ExpiresIn             uint   `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn uint   `json:"refresh_token_expires_in"`
}

type RefreshTokenRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponseDTO struct {
	Token                 string `json:"token"`
	ExpiresIn             uint   `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn uint   `json:"refresh_token_expires_in"`
}
