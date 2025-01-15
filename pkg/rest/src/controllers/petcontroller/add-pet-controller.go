package petcontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// InsertNewPetWithImage
//
//	@Summary	adds a new pet with image
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Param		pet	formData	string	true	"pet"
//	@Success	200	{object}	models.Response
//	@Failure	422	{object}	models.Error
//	@Failure	400	{object}	models.Error
//	@Failure	500	{object}	models.Error
//	@Router		/pet/insert-pet [post]
//	@Security	BearerAuth
func (pc *PetController) InsertNewPetWithImage(ctx *gin.Context) {
	var petReq petmodel.Pet
	var petProfilePic petmodel.ProfilePicture

	petData := ctx.PostForm("pet")
	// unmarshal pet data
	if err := json.Unmarshal([]byte(petData), &petReq); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusUnprocessableEntity, "Status Unprocessable Entity")
		return
	}
	// validate pet data
	if err := petReq.Validate(); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Status Bad Request")
		return
	}

	petReq.ID = uuid.New().String()
	userId := ctx.GetString("user_id")
	// check if user id is empty
	if userId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Internal server error")
		return
	}

	// check if pet already exists
	pet, err := pc.petService.CheckPetExists(petReq, userId)
	// throw error if some error occured
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	// if pet already exists, throw error
	if pet {
		utils.HTTPErrorHandler(ctx, errors.New("pet already exists"), http.StatusBadRequest, "Status Bad Request")
		return
	}
	// if pet does not exists, insert to db
	petImage, _ := ctx.FormFile("pet_image")
	// if profile image is not nil, then insert to imagekit and insert to db
	if petImage != nil && !pet {
		petEncodedImage, err := pc.convertImageToBase64(petImage)
		if err != nil {
			utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
			return
		}
		// Insert to imagekit
		petPicResponse, err := pc.petService.InsertNewPetImage(petmodel.PetProfilePicture{UserId: userId, PetId: petReq.ID, ImageBytes: petEncodedImage, FileName: petImage.Filename})
		if err != nil {
			utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
			return
		}
		// Insert to variable
		petProfilePic.FileId = petPicResponse.Data.FileId
		petProfilePic.Name = petPicResponse.Data.Name
		petProfilePic.Url = petPicResponse.Data.Url
		petProfilePic.ThumbnailUrl = petPicResponse.Data.ThumbnailUrl

		petReq.ProfilePicture = petProfilePic
		// Insert to db
		petRes, err := pc.petService.AddNewPet(petReq, userId)
		if err != nil {
			utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
			return
		}
		utils.HTTPResponseHandler(ctx, petRes, http.StatusOK, "Pet added successfully")
		return
	}
	// Insert to db
	petRes, err := pc.petService.AddNewPet(petReq, userId)
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.HTTPResponseHandler(ctx, petRes, http.StatusOK, "Pet added successfully")
}
