package usermodel

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AboutUser struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	Email  string `json:"email" validate:"required,email"`
	DOB    string `json:"dob" validate:"required"`
	Gender string `json:"gender" validate:"required,oneof=male female other"`
}

func (a *AboutUser) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)
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
			return fmt.Errorf("%s", strings.Join(errMsgs, ", "))
		}
		return err
	}
	return nil
}
