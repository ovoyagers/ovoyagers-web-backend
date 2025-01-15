package usercontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// UpdateUser updates a user's profile
//
//	@Summary		Update User
//	@Tags			users
//	@Description	Update a user's profile
//	@Accept			json
//	@Produce		json
//	@Param			user	body		usermodel.UpdateUser	true	"User object"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Error
//	@Failure		401		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/user/update-user [put]
//	@Security		BearerAuth
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		utils.HTTPErrorHandler(ctx, errors.New("no token provided"), http.StatusUnauthorized, "Unauthorized access")
		return
	}

	var user usermodel.UpdateUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Bad request")
		return
	}

	updatedUser, err := uc.userService.UpdateUser(user, userID.(string))
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}
