package utils

import (
	"avengers-clinic/model/dto/json"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validated(s interface{}) []json.ValidationField {
	validate := validator.New(validator.WithRequiredStructEnabled())

	// register custom validate
	registerValidation(validate)

	var errors []json.ValidationField
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldError := json.ValidationField{
				FieldName: err.Field(),
				Message: getErrorMessage(err),
			}
			errors = append(errors, fieldError)
		}
	}
	return errors
}

func getErrorMessage(err validator.FieldError) string {
	messages := map[string]string{
		"boolean": "Field must be boolean",
		"datetime": "Invalid date",
		"email": "Email is not valid",
		"enum": "Field must be ["+err.Param()+"]",
		"number": "Field must be number",
		"numeric": "Field must be numeric",
		"required": "Field is required",
		"uuid": "Invalid uuid",
		"uuid3": "Invalid uuid",
		"uuid4": "Invalid uuid",
		"lt":       "weekend tutup (>5)",
		"gt":       "weekend tutup (<=0)",
	}

	for key, message := range messages {
		if err.Tag() == key {
			return message
		}
	}
	
	return ""
}

func registerValidation(validate *validator.Validate) {
	validate.RegisterValidation("enum", enumValidator)
	validate.RegisterValidation("regex", regexValidate)
}

func enumValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	params := strings.Split(fl.Param(), " ")
	
	isValid := false
	for _, param := range params {
		if value == param {
			isValid = true
			break
		}
	}
	return isValid
}

func regexValidate(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	pattern := fl.Param()

	matched, _ := regexp.MatchString(pattern, value)
	return matched
}