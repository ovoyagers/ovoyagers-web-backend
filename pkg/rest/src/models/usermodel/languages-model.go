package usermodel

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Languages struct {
	PreferredLanguages []string `json:"preferredLanguages" validate:"required,languageCheck"`
	NativeLanguages    []string `json:"nativeLanguages" validate:"required,languageCheck,differentLanguages"`
}

var validLanguages = []string{"English", "Spanish", "French", "German", "Telugu", "Tamil", "Hindi", "Gujarati", "Marathi"}

func languageCheck(fl validator.FieldLevel) bool {
	languages := fl.Field().Interface().([]string)
	for _, lang := range languages {
		if !isValidLanguage(lang) {
			return false
		}
	}
	return true
}

func isValidLanguage(language string) bool {
	for _, validLanguage := range validLanguages {
		if strings.EqualFold(language, validLanguage) {
			return true
		}
	}
	return false
}

func differentLanguages(fl validator.FieldLevel) bool {
	languages := fl.Field().Interface().([]string)
	otherLanguages := fl.Parent().FieldByName("PreferredLanguages").Interface().([]string)

	for _, lang := range languages {
		for _, otherLang := range otherLanguages {
			if strings.EqualFold(lang, otherLang) {
				return false
			}
		}
	}
	return true
}

var customLanguagesErrorMessages = map[string]string{
	"PreferredLanguages.required":        "Preferred languages are required",
	"PreferredLanguages.languageCheck":   "Preferred languages must be one of the allowed languages",
	"NativeLanguages.required":           "Native languages are required",
	"NativeLanguages.languageCheck":      "Native languages must be one of the allowed languages",
	"NativeLanguages.differentLanguages": "Native languages must not be in preferred languages",
}

func (l *Languages) Validate() error {
	validate := validator.New()

	// Register custom validations
	err := validate.RegisterValidation("languageCheck", languageCheck)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("differentLanguages", differentLanguages)
	if err != nil {
		return err
	}

	err = validate.Struct(l)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []string
			for _, validationErr := range validationErrors {
				field := strings.Split(validationErr.StructNamespace(), ".")[1]
				tag := validationErr.Tag()
				key := field + "." + tag
				fieldKey := strings.Split(key, ".")[1]
				if msg, found := customLanguagesErrorMessages[key]; found {
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
