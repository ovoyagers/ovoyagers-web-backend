package petmodel

import (
	"time"
)

type ProfilePicture struct {
	FileId       string `json:"fileId" default:""`
	Url          string `json:"url" default:""`
	Name         string `json:"name" default:""`
	ThumbnailUrl string `json:"thumbnailUrl" default:""`
}

// Pet struct definition
type Pet struct {
	ID             string         `json:"id"`
	Name           string         `json:"name" validate:"required" example:"Bruno"`
	Dob            time.Time      `json:"dob" validate:"required" example:"2024-01-01T00:00:00Z"`
	Gender         string         `json:"gender" validate:"required,oneof=male female other" example:"male"`
	Kind           string         `json:"kind" validate:"required" example:"dog"`
	Breed          string         `json:"breed" validate:"required" example:"golden-retriever"`
	Weight         float64        `json:"weight" example:"10.5"`
	IsPrimary      bool           `json:"isPrimary" example:"false"`
	ProfilePicture ProfilePicture `json:"profilePicture,omitempty"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

// Validate method for the Pet struct
func (p *Pet) Validate() error {
	// validate struct
	return validateStruct(p)
}
