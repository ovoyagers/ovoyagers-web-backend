package petcontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
	"github.com/sirupsen/logrus"
)

// DeletePet deletes a pet
//
//	@Summary	deletes a pet with the given id
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Param		petId	path		string	true	"pet id"
//	@Param		fileId	query		string	true	"file id"
//	@Success	200		{object}	models.Response
//	@Failure	400		{object}	models.Error
//	@Failure	500		{object}	models.Error
//	@Router		/pet/delete-pet/{petId} [delete]
//	@Security	BearerAuth
func (pc *PetController) DeletePet(ctx *gin.Context) {
	var finalResp *api.Response
	var err error
	// get pet id from url
	petId := ctx.Param("petId")
	if petId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("pet id not found"), http.StatusBadRequest, "Status Bad Request")
		return
	}
	fileId := ctx.Query("fileId")
	if fileId != "" {
		finalResp, err = pc.petService.DeletePetImage(fileId)
		if err != nil {
			utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Status Internal Server Error")
			return
		}
	}

	userID := ctx.GetString("user_id")
	if userID == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	resp, err := pc.petService.DeletePetFolder(userID, petId)
	if err != nil && err.Error() != "Not Found" {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}
	logrus.Info(resp)
	// delete pet
	err = pc.petService.DeletePet(petId)
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}
	utils.HTTPResponseHandler(ctx, finalResp, http.StatusOK, "Pet deleted successfully")
}

// DeletePetImage deletes a pet image from the ImageKit service.
//
// @Summary	Deletes a pet image from the ImageKit service.
// @Tags		pet
// @Accept		json
// @Produce	json
// @Param		petId	path		string	true	"pet id"
// @Param		fileId	query		string	true	"file id"
// @Success	200		{object}	models.Response
// @Failure	400		{object}	models.Error
// @Failure	500		{object}	models.Error
// @Router		/pet/delete-pet-image/{petId} [delete]
// @Security	BearerAuth
func (pc *PetController) DeletePetImage(ctx *gin.Context) {
	fileId := ctx.Query("fileId")
	if fileId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("file id not found"), http.StatusBadRequest, "Status Bad Request")
		return
	}
	petId := ctx.Param("petId")
	if petId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("pet id not found"), http.StatusBadRequest, "Status Bad Request")
		return
	}

	if err := pc.petService.DeletePetProfilePicture(fileId, petId); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	utils.HTTPResponseHandler(ctx, nil, http.StatusOK, "Pet image deleted successfully")
}
