package invitemodel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var customErrorMessages = map[string]string{
	"email.required": "Email is required",
	"email.email":    "Email is invalid",
	"status.oneof":   "Status is invalid",
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
