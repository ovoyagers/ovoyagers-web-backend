package notificationscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	notificationmodel "github.com/petmeds24/backend/pkg/rest/src/models/notification-model"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// SendNotification sends a push notification to a user using the FCM Token
//
//	@Summary		Sends a push notification to a user using the FCM Token
//	@Description	Sends a push notification to a user using the FCM Token
//	@Tags			notifications
//	@Accept			json
//	@Produce		json
//	@Param			notification	body		notificationmodel.SendNotificationRequest	true	"Notification"
//	@Success		200				{object}	models.Response
//	@Failure		400				{object}	models.Error
//	@Failure		422				{object}	models.Error
//	@Failure		401				{object}	models.Error
//	@Failure		404				{object}	models.Error
//	@Failure		500				{object}	models.Error
//	@Router			/notifications/send [post]
//	@security		BearerAuth
func (c *NotificationsController) SendNotification(ctx *gin.Context) {
	var request notificationmodel.SendNotificationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Bad request")
		return
	}
	if err := request.Validate(); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusUnprocessableEntity, "Unprocessable entity")
		return
	}

	if err := c.pushSvc.SendToSingleDevice(request); err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.HTTPResponseHandler(ctx, nil, http.StatusOK, "Notification sent successfully")
}
