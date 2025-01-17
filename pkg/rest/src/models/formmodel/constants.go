package formmodel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Custom error messages for validation errors
var customErrorMessages = map[string]string{
	"Mobile.required":   "phone is required",
	"Mobile.e164":       "phone must be in E.164 format",
	"Fullname.min":      "full name must be at least 2 characters",
	"Fullname.max":      "full name must be at most 100 characters",
	"Email.required":    "email is required",
	"Email.email":       "email must be a valid email address",
	"Message.required":  "message is required",
	"Category.required": "category is required",
	"Category.oneof":    "category must be one of 'contact', 'hotel', or 'flights'",
}

func validateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []string
			for _, validationErr := range validationErrors {
				field := strings.Split(validationErr.StructNamespace(), ".")[1]
				tag := validationErr.Tag()
				key := field + "." + tag
				fieldKey := strings.Split(key, ".")[1]
				if msg, found := customErrorMessages[key]; found {
					errMsgs = append(errMsgs, msg)
				} else {
					errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", strings.ToLower(fieldKey)))
				}
			}
			return errors.New(strings.Join(errMsgs, ", "))
		}
		return err
	}
	return nil
}
