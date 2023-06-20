package validator

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateStruct(s interface{}) error {
	if validate == nil {
		validate = validator.New()
	}

	return validate.Struct(s)
}
