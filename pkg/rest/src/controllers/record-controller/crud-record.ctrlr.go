package recordcontroller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	recordmodel "github.com/petmeds24/backend/pkg/rest/src/models/record-model"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
	imgUtils "github.com/petmeds24/backend/pkg/rest/src/utils/images"
)

// InsertMedicalRecords adds a new medical record
//
//	@Summary		adds a new medical record
//	@Description	adds a new medical record
//	@Tags			records
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			records		formData	file	true	"records"
//	@Param			description	formData	string	true	"description"
//	@Success		200			{object}	models.Response
//	@Failure		400			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/record/insert-medical-records [post]
//	@Security		BearerAuth
func (rc *RecordController) InsertMedicalRecords(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "status bad request")
		return
	}

	records := form.File["records"]
	description := form.Value["description"][0]
	petId := form.Value["petId"][0]

	// validate request
	if len(records) == 0 {
		// fmt.Println("No records found")
		utils.HTTPErrorHandler(c, fmt.Errorf("no records found"), http.StatusBadRequest, "no records found")
		return
	}
	var converted []*models.ImgMetaData
	for _, record := range records {
		recordByte, err := imgUtils.ConvertImageToBase64(record, record.Filename)
		if err != nil {
			utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error converting image to base64")
			return
		}
		converted = append(converted, recordByte)
	}
	var recs []recordmodel.Image

	for _, record := range converted {
		newrecord := recordmodel.Image{
			ImageBytes: record.Avatar,
			FileName:   record.Filename,
		}
		recs = append(recs, newrecord)
	}
	// get image metadata
	data := recordmodel.MedicalRecordImageMetadata{
		UserId: c.GetString("user_id"),
		PetId:  petId,
		Image:  recs,
	}
	// insert records to imageKit
	res, err := rc.recordSvc.InsertRecordImages(data)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error inserting records to imageKit")
		return
	}
	var imgRefs []recordmodel.PetRecordImage
	for _, record := range res {
		imgRef := recordmodel.PetRecordImage{
			FileId:   record.Data.FileId,
			ImageUrl: record.Data.Url,
		}
		imgRefs = append(imgRefs, imgRef)
	}
	recordsData := recordmodel.RecordData{
		UserId:      c.GetString("user_id"),
		PetId:       petId,
		Description: description,
		PetRecords:  imgRefs,
	}

	// insert records to database
	finalRes, err := rc.recordSvc.InsertMedicalRecordDao(recordsData)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error inserting records to database")
		return
	}
	// convert string to JSON
	var petRecords []recordmodel.PetRecordImage
	err = utils.StringToJSON(finalRes["pet_records"].(string), &petRecords)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error converting string to JSON")
		return
	}
	finalRes["pet_records"] = petRecords
	// return response
	utils.HTTPResponseHandler(c, finalRes, http.StatusOK, "success")
}

// GetMedicalRecordsByPetId retrieves the medical record of a pet given its id.
//
//	@Summary		retrieves the medical record of a pet given its id.
//	@Description	retrieves the medical record of a pet given its id.
//	@Tags			records
//	@Accept			application/json
//	@Produce		json
//	@Param			petId	path		string	true	"pet id"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/record/get-medical-records/{petId} [get]
//	@Security		BearerAuth
func (rc *RecordController) GetMedicalRecordsByPetId(c *gin.Context) {
	petId := c.Param("petId")
	res, err := rc.recordSvc.GetMedicalRecordsByPetId(petId)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error getting medical records")
		return
	}
	utils.HTTPResponseHandler(c, res, http.StatusOK, "success")
}

// DeleteMedicalRecordsById deletes the medical record of a pet given its id.
//
//	@Summary		deletes the medical record of a pet given its id.
//	@Description	deletes the medical record of a pet given its id.
//	@Tags			records
//	@Accept			application/json
//	@Produce		json
//	@Param			recordId	path		string										true	"record id"
//	@Param			imageIds	body		recordmodel.DeletePetMedicalRecordRequest	true	"image ids"
//	@Success		200			{object}	models.Response
//	@Failure		400			{object}	models.Error
//	@Failure		422			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/record/delete-medical-record/{recordId} [post]
//	@Security		BearerAuth
func (rc *RecordController) DeleteMedicalRecordById(c *gin.Context) {
	recordId := c.Param("recordId")
	if recordId == "" {
		utils.HTTPErrorHandler(c, errors.New("record ID is required"), http.StatusUnprocessableEntity, "missing record ID")
		return
	}
	var petMedicalRecordRequest recordmodel.DeletePetMedicalRecordRequest
	if err := c.ShouldBindJSON(&petMedicalRecordRequest); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusUnprocessableEntity, "Status Unprocessable Entity")
		return
	}
	// delete medical record from database
	err := rc.recordSvc.DeleteRecordImages(petMedicalRecordRequest.FileIds, recordId)
	// delete medical record from imageKit
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error deleting medical record")
		return
	}
	// return response
	utils.HTTPResponseHandler(c, nil, http.StatusOK, "success")
}

// GetMedicalRecordById retrieves the medical record of a pet given its id.
//
//	@Summary		retrieves the medical record of a pet given its id.
//	@Description	retrieves the medical record of a pet given its id.
//	@Tags			records
//	@Accept			application/json
//	@Produce		json
//	@Param			medicalRecordId	path		string	true	"medical record id"
//	@Success		200				{object}	models.Response
//	@Failure		400				{object}	models.Error
//	@Failure		500				{object}	models.Error
//	@Router			/record/get-medical-record/{medicalRecordId} [get]
//	@Security		BearerAuth
func (rc *RecordController) GetMedicalRecordById(c *gin.Context) {
	medicalRecordId := c.Param("medicalRecordId")
	if medicalRecordId == "" {
		utils.HTTPErrorHandler(c, errors.New("medical record ID is required"), http.StatusUnprocessableEntity, "missing medical record ID")
		return
	}
	res, err := rc.recordSvc.GetMedicalRecordByRecordId(medicalRecordId)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "error getting medical record")
		return
	}
	utils.HTTPResponseHandler(c, res, http.StatusOK, "success")
}
