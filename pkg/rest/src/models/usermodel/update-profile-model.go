package usermodel

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UpdateUser struct {
	Fullname    string `json:"fullname" validate:"required,min=3,max=64" example:"Pecol"`
	Phone       string `json:"phone" validate:"required,min=10,max=10,numeric" example:"6302068026"`
	CountryCode string `json:"countryCode" validate:"required,iso3166_1_alpha2" example:"IN"`
	Age         uint8  `json:"age" default:"0" validate:"gte=0,lte=130" example:"22"`
	Gender      string `json:"gender" default:"" validate:"oneof=male female other" example:"male"`
}

func (u *UpdateUser) Validate() error {
	// create validator instance
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		var errorMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			if msg, found := customErrorMessages[err.Field()]; found {
				errorMsgs = append(errorMsgs, msg)
			} else {
				errorMsgs = append(errorMsgs, fmt.Sprintf("%s is invalid", strings.ToLower(err.Field())))
			}
		}
		return fmt.Errorf("%s", strings.Join(errorMsgs, ", "))
	}
	return nil
}
