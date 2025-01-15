package recordmodel

type DeletePetMedicalRecordRequest struct {
	FileIds []string `json:"fileIds" binding:"required"`
}
