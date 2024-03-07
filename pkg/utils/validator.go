package utils

import (
	"avengers-clinic/model/dto/json"

	"github.com/go-playground/validator/v10"
)

func Validated(s interface{}) []json.ValidationField {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var errors []json.ValidationField
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldError := json.ValidationField{
				FieldName: err.Field(),
				Message: getErrorMessage(err.Tag()),
			}
			errors = append(errors, fieldError)
		}
	}
	return errors
}

func getErrorMessage(tag string) string {
	messages := map[string]string{
		"boolean": "Field must be boolean",
		"datetime": "Field must be date",
		"email": "Email is not valid",
		"number": "Field must be number",
		"numeric": "Field must be numeric",
		"required": "Field is required",
		"uuid": "Invalid uuid",
		"uuid3": "Invalid uuid",
		"uuid4": "Invalid uuid",
	}

	for key, message := range messages {
		if tag == key {
			return message
		}
	}

	return ""
}