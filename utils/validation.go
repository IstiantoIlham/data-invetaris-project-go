package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateInput(input interface{}) (map[string]string, error) {
	err := validate.Struct(input)
	if err == nil {
		return nil, nil
	}
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		tag := err.ActualTag()
		errors[fieldName] = fmt.Sprintf("Validation failed for '%s' (%s)", fieldName, tag)
	}
	return errors, fmt.Errorf("validation failed")
}
