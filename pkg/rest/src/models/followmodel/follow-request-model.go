package followmodel

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FollowRequest struct {
	FriendUsername string `json:"friend_username" validate:"required"`
}

func (fr *FollowRequest) Validate() error {
	// create validator instance
	validate := validator.New()
	// validate struct
	err := validate.Struct(fr)
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
