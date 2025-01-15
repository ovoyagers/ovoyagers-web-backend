package recordmodel

import "mime/multipart"

type MedicalRecordRequest struct {
	RecordFiles []*multipart.FileHeader `form:"recordFiles" binding:"required,min=1"`
	Description string                  `form:"description" binding:"required"`
}

type Image struct {
	ImageBytes string `json:"imageBytes"`
	FileName   string `json:"fileName"`
}

type MedicalRecordImageMetadata struct {
	UserId  string  `json:"userId"`
	PetId   string  `json:"petId"`
	Image   []Image `json:"image"`
	ImageId string  `json:"imageId"`
}

type PetRecordImage struct {
	FileId   string `json:"fileId"`
	ImageUrl string `json:"imageUrl"`
}

type RecordData struct {
	PetRecords  []PetRecordImage `json:"petRecords"`
	UserId      string           `json:"userId"`
	PetId       string           `json:"petId"`
	Description string           `json:"description"`
}
