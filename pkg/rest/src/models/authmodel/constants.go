package authmodel

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Custom error messages for validation errors
var customErrorMessages = map[string]string{
	"CountryCode.iso3166_1_alpha2": "countryCode must be a valid ISO 3166-1 alpha-2 country code",
	"CountryCode.required":         "countryCode is required",
	"Phone.required":               "phone is required",
	"Phone.e164":                   "phone must be in E.164 format",
	"Fullname.required":            "full name is required",
	"Fullname.min":                 "full name must be at least 2 characters",
	"Fullname.max":                 "full name must be at most 100 characters",
	"Email.required":               "email is required",
	"Email.email":                  "email must be a valid email address",
	"Email.emaildomain":            "email must be a valid ovoyagers.com email address",
	"DOB.required":                 "dob is required",
	"Age.gte":                      "age must be at least 13",
	"Age.lte":                      "age must be at most 130",
	"Gender.required":              "gender is required",
	"Gender.oneof":                 "gender must be one of 'male', 'female', or 'other'",
	"Code.required":                "code is required",
	"Code.min":                     "code must be at least 6 characters",
	"Code.max":                     "code must be at most 6 characters",
	"Password.required":            "password is required",
	"Password.min":                 "password must be at least 8 characters",
	"Password.max":                 "password must be at most 64 characters",
	"Password.alphanum":            "password must only contain letters and numbers",
}

func validateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.RegisterValidation("emaildomain", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		domain := strings.Split(email, "@")[1]
		return domain == "ovoyagers.com"
	})

	if err != nil {
		return err
	}

	err = validate.RegisterValidation("pswd", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		// Check regex with at least 8 characters, 1 uppercase, 1 lowercase, 1 number, and 1 special character
		return regexp.MustCompile(`^[a-zA-Z\d@$!%*?&]{8,}$`).MatchString(password) &&
			regexp.MustCompile(`[a-z]`).MatchString(password) &&
			regexp.MustCompile(`[A-Z]`).MatchString(password) &&
			regexp.MustCompile(`\d`).MatchString(password) &&
			regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)
	})
	if err != nil {
		return err
	}
	if err := validate.Struct(s); err != nil {
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
