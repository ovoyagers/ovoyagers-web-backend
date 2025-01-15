package followcontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/models/followmodel"
	"github.com/petmeds24/backend/pkg/rest/src/services/followservice"
	log "github.com/sirupsen/logrus"
)

type FollowController struct {
	followService *followservice.FollowService
}

func NewFollowController(ctx context.Context) *FollowController {
	return &FollowController{
		followService: followservice.NewFollowService(ctx),
	}
}

// CreateFollowRequest creates a new follow request between two users
//
//	@Summary		Create Follow Request
//	@Description	Create Follow Request
//	@Tags			follow
//	@Accept			json
//	@Produce		json
//	@Param			followingUser	body		followmodel.FollowRequest	true	"Following User"
//	@Success		200				{object}	models.Response
//	@Failure		422				{object}	models.Error
//	@Failure		400				{object}	models.Error
//	@Failuer		401             {object}  models.Error
//	@Failure		404	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/follow/create [post]
//	@Security		BearerAuth
func (fc *FollowController) CreateFollowRequest(c *gin.Context) {
	// fetch user data from request body
	var followRequest followmodel.FollowRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		log.Error(err)
		c.JSON(http.StatusUnprocessableEntity, models.Error{
			Message:    err.Error(),
			Error:      "Status Unprocessable Entity",
			Status:     "error",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}

	// validate following user data
	if err := followRequest.Validate(); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    err.Error(),
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// check if username exists
	userid := c.GetString("user_id")

	if userid == "" {
		log.Error("user_id not found in context")
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    "user_id not found in context",
			Error:      "Status Unauthorized",
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	// create follow request
	data, err := fc.followService.CreateFollowRequest(followRequest.FriendUsername, userid)

	// TODO: CALL NOTIFICATION SERVICE TO SEND NOTIFICATION TO FRIEND

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{
			Message:    err.Error(),
			Error:      "Status Internal Server Error",
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message:    "Follow request created successfully",
		Data:       data,
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}

// AcceptFollowRequest accepts a follow request
//
//	@Summary		Accept Follow Request
//	@Description	Accept Follow Request
//	@Tags			follow
//	@Accept			json
//	@Produce		json
//	@Param			followingUser	body		followmodel.FollowRequest	true	"Following User"
//	@Success		200				{object}	models.Response
//	@Failure		422				{object}	models.Error
//	@Failure		404				{object}	models.Error
//	@Failure		400				{object}	models.Error
//	@Failuer		401             {object}  models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/follow/accept [post]
//	@Security		BearerAuth
func (fc *FollowController) AcceptFollowRequest(c *gin.Context) {
	// fetch user data from request body
	var followRequest followmodel.FollowRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		log.Error(err)
		c.JSON(http.StatusUnprocessableEntity, models.Error{
			Message:    err.Error(),
			Error:      "Status Unprocessable Entity",
			Status:     "error",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}

	// validate following user data
	if err := followRequest.Validate(); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    err.Error(),
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// check if username exists
	userid := c.GetString("user_id")

	if userid == "" {
		log.Error("user_id not found in context")
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    "user_id not found in context",
			Error:      "Status Unauthorized",
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	// Accept follow request
	data, err := fc.followService.AcceptFollowRequest(followRequest.FriendUsername, userid)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{
			Message:    err.Error(),
			Error:      "Status Internal Server Error",
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	// TODO: CALL NOTIFICATION SERVICE TO SEND NOTIFICATION TO FRIEND

	// return response
	c.JSON(http.StatusOK, models.Response{
		Message:    "Follow request accepted successfully",
		Data:       data,
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}

// CancelFollowRequest cancels a follow request
//
//	@Summary		Cancel Follow Request
//	@Description	Cancel Follow Request
//	@Tags			follow
//	@Accept			json
//	@Produce		json
//	@Param			followingUser	body		followmodel.FollowRequest	true	"Following User"
//	@Success		200				{object}	models.Response
//	@Failure		422				{object}	models.Error
//	@Failure		400				{object}	models.Error
//	@Failuer		401             {object}  models.Error
//	@Failure		404	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/follow/cancel [post]
//	@Security		BearerAuth
func (fc *FollowController) CancelFollowRequest(c *gin.Context) {
	// fetch user data from request body
	var followRequest followmodel.FollowRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		log.Error(err)
		c.JSON(http.StatusUnprocessableEntity, models.Error{
			Message:    err.Error(),
			Error:      "Status Unprocessable Entity",
			Status:     "error",
			StatusCode: http.StatusUnprocessableEntity,
		})

		return
	}

	// validate following user data
	if err := followRequest.Validate(); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    err.Error(),
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// check if username exists
	userid := c.GetString("user_id")

	if userid == "" {
		log.Error("user_id not found in context")
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    "user_id not found in context",
			Error:      "Status Unauthorized",
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	// Cancel follow request
	err := fc.followService.CancelFollowRequest(followRequest.FriendUsername, userid)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{
			Message:    err.Error(),
			Error:      "Status Internal Server Error",
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message:    "Follow request canceled successfully",
		Data:       nil,
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}
