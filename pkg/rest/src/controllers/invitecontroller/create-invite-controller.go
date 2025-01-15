package invitecontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/petmeds24/backend/pkg/rest/src/models/invitemodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// SendInvite invite to user
//
//	@Summary		invite to user
//	@Description	invite to user
//	@Tags			invite
//	@Accept			json
//	@Produce		json
//	@Param			invite	body		invitemodel.InviteUser	true	"invite"
//	@Success		200		{object}	models.Response
//	@Failure		422		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/invite/send-invite [post]
//	@Security		BearerAuth
func (ic *InviteController) SendInvite(ctx *gin.Context) {
	// get user id from context
	userId := ctx.GetString("user_id")
	if userId == "" {
		utils.HTTPErrorHandler(ctx, errors.New("user id not found"), http.StatusInternalServerError, "Internal server error")
		return
	}

	// bind invite user data from request body
	var inviteUser invitemodel.InviteUser
	if err := ctx.ShouldBindJSON(&inviteUser); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Status Bad Request")
		return
	}

	// validate invite user
	if err := inviteUser.Validate(); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Status Bad Request")
		return
	}

	// add invite to inviteuser
	inviteUser.Id = uuid.New().String()

	// send invite to user
	if err := ic.inviteSvc.SendInvite(inviteUser, userId); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}

	// send email to user
	if err := ic.inviteSvc.SendInviteEmail(inviteUser.Email, inviteUser.Id); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.HTTPResponseHandler(ctx, nil, http.StatusOK, "Invite sent successfully")
}
