package authmodel

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProfilePicture struct {
	FileId       string `json:"fileId" default:""`
	Url          string `json:"url" default:""`
	Name         string `json:"name" default:""`
	ThumbnailUrl string `json:"thumbnailUrl" default:""`
}

// Pet struct definition
type User struct {
	ID             string         `json:"id,omitempty"`
	Fullname       string         `json:"fullname" validate:"required,min=2,max=100"`
	CountryCode    string         `json:"countryCode" validate:"required,iso3166_1_alpha2"`
	Phone          string         `json:"phone" validate:"required,e164"`
	Email          string         `json:"email" default:"" validate:"email,required"`
	Password       string         `json:"password" validate:"required,min=8,max=32"`
	DOB            string         `json:"dob" default:""`
	Age            uint8          `json:"age" default:"0" validate:"gte=0,lte=130"`
	Gender         string         `json:"gender" default:"" validate:"oneof=male female other"`
	ProfilePicture ProfilePicture `json:"profilePicture,omitempty"`
	IsVerified     bool           `json:"isVerified" default:"false"`
	Active         bool           `json:"active" default:"false"`
	Coins          uint64         `json:"coins" default:"0"`
	OTP            string         `json:"otp" default:""`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

type CheckUser struct {
	Email string `json:"email,omitempty" validate:"required,email"`
	Id    string `json:"id,omitempty"`
}

// Validate method for the Pet struct
func (u *User) Validate() error {
	// create validator instance
	validate := validator.New()
	// validate struct
	err := validate.Struct(u)
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

					errorMsgs = append(errorMsgs, strings.ToLower(fieldKey)+" is invalid")
				}
			}
			return errors.New(strings.Join(errorMsgs, ", "))
		}
		return err
	}
	return nil
}

// StructToMap converts the Pet struct to a map
func (u *User) StructToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             u.ID,
		"active":         u.Active,
		"countryCode":    u.CountryCode,
		"phone":          u.Phone,
		"fullname":       u.Fullname,
		"email":          u.Email,
		"age":            u.Age,
		"gender":         u.Gender,
		"profilePicture": u.ProfilePicture,
		"isVerified":     u.IsVerified,
		"otp":            u.OTP,
		"password":       u.Password,
		"dob":            u.DOB,
		"coins":          u.Coins,
		"createdAt":      u.CreatedAt,
		"updatedAt":      u.UpdatedAt,
	}
}
