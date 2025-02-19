package utils

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	var validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil

}
