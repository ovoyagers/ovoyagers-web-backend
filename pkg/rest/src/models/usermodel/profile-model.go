package usermodel

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserProfile struct {
	UserName       string `json:"username"`
	Name           string `json:"name"`
	DOB            string `json:"dob"`
	Gender         string `json:"gender" validate:"oneof=male female other"`
	Profession     string `json:"profession"`
	TextBio        string `json:"bio" validate:"max=500"`
	AudioBio       string `json:"audioBio"`
	ProfilePicture string `json:"profilePicture"`
	IsBioAudio     bool   `json:"isBioAudio"`
}

func (up *UserProfile) Validate() error {
	// create validator instance
	validate := validator.New()
	// validate struct
	err := validate.Struct(up)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMsgs []string
			for _, validationErr := range validationErrors {
				field := strings.Split(validationErr.StructNamespace(), ".")[1]
				tag := validationErr.Tag()
				key := field + "." + tag
				fieldKey := strings.Split(key, ".")[1]
				if msg, found := customErrorMessages[key]; found {
					errorMsgs = append(errorMsgs, msg)
				} else {
					errorMsgs = append(errorMsgs, fmt.Sprintf("%s is invalid", strings.ToLower(fieldKey)))
				}
			}
			return fmt.Errorf("%s", strings.Join(errorMsgs, ", "))
		}
		return err
	}
	return nil
}

func (up *UserProfile) ConvertStructToMap() map[string]interface{} {
	return map[string]interface{}{
		"username":       up.UserName,
		"name":           up.Name,
		"gender":         up.Gender,
		"profession":     up.Profession,
		"textBio":        up.TextBio,
		"audioBio":       up.AudioBio,
		"profilePicture": up.ProfilePicture,
		"isBioAudio":     up.IsBioAudio,
	}
}
