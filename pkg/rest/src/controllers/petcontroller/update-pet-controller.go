package petcontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// UpdatePet updates the  pet of the user based on the petId
//
//	@Summary	updates the  pet of the user based on the petId
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string			true	"pet id"
//	@Param		pet	body		petmodel.Pet	true	"pet details"
//	@Success	200	{object}	petmodel.Pet
//	@Failure	400	{object}	models.Error
//	@Failure	500	{object}	models.Error
//	@Router		/pet/update-pet [put]
//	@Security	BearerAuth
func (pc *PetController) UpdatePet(ctx *gin.Context) {
	userid := ctx.GetString("user_id")
	if userid == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Internal server error")
		return
	}
	petId := ctx.Param("id")
	if petId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("pet id is not provided, please check"), http.StatusBadRequest, "Bad request")
		return
	}
	var pet petmodel.Pet
	if err := ctx.ShouldBindJSON(&pet); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Bad request")
		return
	}

	if err := pc.petService.UpdatePet(pet, petId); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.HTTPResponseHandler(ctx, nil, http.StatusOK, "Pet updated successfully")
}

// UpdatePetImage updates the  pet image of the user based on the petId
//
//	@Summary	updates the  pet image of the user based on the petId
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string	true	"pet id"
//	@Param		avatar	formData	file	true	"Pet Image"
//	@Success	200		{object}	petmodel.Pet
//	@Failure	400		{object}	models.Error
//	@Failure	500		{object}	models.Error
//	@Router		/pet/update-pet-image [put]
//	@Security	BearerAuth
func (pc *PetController) UpdatePetImage(ctx *gin.Context) {
	userid := ctx.GetString("user_id")
	if userid == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Internal server error")
		return
	}
	petId := ctx.Param("id")
	if petId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("pet id is not provided, please check"), http.StatusBadRequest, "Bad request")
		return
	}

	file, err := ctx.FormFile("avatar")
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Bad request")
		return
	}

	petEncodedImage, err := pc.convertImageToBase64(file)
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}

	petImg := petmodel.PetProfilePicture{UserId: userid, PetId: petId, ImageBytes: petEncodedImage, FileName: file.Filename}

	// update pet image
	if err := pc.petService.UpdatePetImage(petImg); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.HTTPResponseHandler(ctx, nil, http.StatusOK, "Pet image updated successfully")
}
