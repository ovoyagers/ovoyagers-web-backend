package petmodel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Custom error messages for validation errors
var customErrorMessages = map[string]string{
	"Name.required":   "name is required",
	"Name.min":        "name must be at least 2 characters",
	"Name.max":        "name must be at most 100 characters",
	"Kind.required":   "kind is required",
	"Kind.min":        "kind must be at least 2 characters",
	"Kind.max":        "kind must be at most 100 characters",
	"Breed.required":  "breed is required",
	"Breed.min":       "breed must be at least 2 characters",
	"Breed.max":       "breed must be at most 100 characters",
	"Dob.required":    "dob is required",
	"Dob.date":        "dob must be a valid date",
	"Gender.required": "gender is required",
	"Gender.oneof":    "gender must be one of 'male', 'female', or 'other'",
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
					fmt.Println(key, fieldKey)
					errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", strings.ToLower(fieldKey)))
				}
			}
			return errors.New(strings.Join(errMsgs, ", "))
		}
		return err
	}
	return nil
}
