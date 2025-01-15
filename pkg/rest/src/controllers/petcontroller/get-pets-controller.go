package petcontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models/petmodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// GetPrimaryPet returns the primary pet of the user
//
//	@Summary	gets the primary pet of the user
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	petmodel.Pet
//	@Failure	400	{object}	models.Error
//	@Failure	500	{object}	models.Error
//	@Router		/pet/primary-pet [get]
//	@Security	BearerAuth
func (pc *PetController) GetPrimaryPet(ctx *gin.Context) {
	userid := ctx.GetString("user_id")
	if userid == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Internal server error")
		return
	}
	primaryPet, err := pc.petService.GetPrimaryPet(userid)
	if err == utils.ErrNoDataFound {
		utils.HTTPResponseHandler(ctx, nil, http.StatusOK, "Primary pet not found")
		return
	}

	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.HTTPResponseHandler(ctx, primaryPet, http.StatusOK, "Primary pet fetched successfully")
}

// ListPets returns the list of pets of the user
//
//	@Summary	returns the list of pets of the user
//	@Tags		pet
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		petmodel.Pet
//	@Failure	400	{object}	models.Error
//	@Failure	500	{object}	models.Error
//	@Router		/pet/list-pets [get]
//	@Security	BearerAuth
func (pc *PetController) ListPets(ctx *gin.Context) {
	userid := ctx.GetString("user_id")
	if userid == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Internal server error")
		return
	}
	pets, err := pc.petService.ListPets(userid)
	if err == utils.ErrNoDataFound {
		utils.HTTPResponseHandler(ctx, []petmodel.Pet{}, http.StatusOK, "Pets not found")
		return
	}
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.HTTPResponseHandler(ctx, pets, http.StatusOK, "Pets fetched successfully")
}
