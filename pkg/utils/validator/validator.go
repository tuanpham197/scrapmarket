package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateRequestData(req interface{}) []*Error {
	var errors []*Error
	validate = validator.New()
	errValidate := validate.Struct(req)

	if errValidate != nil {
		for _, err := range errValidate.(validator.ValidationErrors) {
			var el Error
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			el.Error = customMessageError(&err)
			errors = append(errors, &el)
		}

		return errors
	}

	return nil
}

func customMessageError(fieldError *validator.FieldError) string {
	var message string = (*fieldError).Error()
	switch (*fieldError).Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", (*fieldError).Field())
	case "email":
		return fmt.Sprintf("Field %s must be a valid email", (*fieldError).Field())
	case "file":
		return fmt.Sprintf("Field %s invalid", (*fieldError).Field())
	case "image":
		return fmt.Sprintf("Field %s invalid", (*fieldError).Field())
	case "required_if":
		return fmt.Sprintf("Field %s invalid", (*fieldError).Field())
	case "max":
		return fmt.Sprintf("Field %s invalid", (*fieldError).Field())
	case "min":
		return fmt.Sprintf("Field %s invalid", (*fieldError).Field())
	}

	return message
}
