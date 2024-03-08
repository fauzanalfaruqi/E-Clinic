package utils

import (
	"avengers-clinic/model/dto/json"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Validated(s interface{}) []json.ValidationField {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("regex", regexValidate)
	var errors []json.ValidationField
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldError := json.ValidationField{
				FieldName: err.Field(),
				Message:   getErrorMessage(err.Tag()),
			}
			errors = append(errors, fieldError)
		}
	}
	return errors
}

func regexValidate(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	pattern := fl.Param()

	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

func getErrorMessage(tag string) string {
	messages := map[string]string{
		"required": "field is required",
		"email":    "email is not valid",
		"string":   "field is not string",
		"number":   "field is not number",
		"lt":       "weekend tutup (>5)",
		"gt":       "weekend tutup (<=0)",
	}

	for key, message := range messages {
		if tag == key {
			return message
		}
	}

	return ""
}
