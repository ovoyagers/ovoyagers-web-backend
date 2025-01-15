package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	"github.com/petmeds24/backend/pkg/rest/src/services/userservice"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
	userService *userservice.UserService
}

func NewUserController(globalCfg *config.GlobalConfig) *UserController {
	return &UserController{
		userService: userservice.NewAuthService(globalCfg),
	}
}

// UpdateAboutUser godoc
//
//	@Summary		Update About User
//	@Description	Update About User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			aboutUser	body		usermodel.AboutUser	true	"About User"
//	@Success		200			{object}	models.Response
//	@Failure		422			{object}	models.Error
//	@Failure		400			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/user/update-about [put]
//	@Security		BearerAuth
func (uc *UserController) UpdateAboutUser(c *gin.Context) {
	var aboutUser usermodel.AboutUser
	if err := c.ShouldBindJSON(&aboutUser); err != nil {
		log.Error(err)
		c.JSON(http.StatusUnprocessableEntity, models.Error{
			Message:    err.Error(),
			Error:      "Status Unprocessable Entity",
			Status:     "error",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}

	// validate user data
	if err := aboutUser.Validate(); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    err.Error(),
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// extract user id from jwt
	userId := c.GetString("user_id")

	// validate user id
	if userId == "" {
		log.Error("user id is empty")
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    "user id is empty",
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// update user
	data, err := uc.userService.UpdateAboutUser(&aboutUser, userId)
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

	if data == nil {
		log.Error("user not found")
		c.JSON(http.StatusNotFound, models.Error{
			Message:    "user not found",
			Error:      "Status Not Found",
			Status:     "error",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpdateLanguages godoc
//
//	@Summary		Update Languages
//	@Description	Update Languages
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			languages	body		usermodel.Languages	true	"Languages"
//	@Success		200			{object}	models.Response
//	@Failure		422			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/user/update-languages [put]
//	@Security		BearerAuth
func (uc *UserController) UpdateLanguages(c *gin.Context) {
	var languages usermodel.Languages
	if err := c.ShouldBindJSON(&languages); err != nil {
		log.Error(err)
		c.JSON(http.StatusUnprocessableEntity, models.Error{
			Message:    err.Error(),
			Error:      "Status Unprocessable Entity",
			Status:     "error",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}

	err := languages.Validate()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    err.Error(),
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// extract user id from jwt
	userId := c.GetString("user_id")

	// validate user id
	if userId == "" {
		log.Error("user id is empty")
		c.JSON(http.StatusBadRequest, models.Error{
			Message:    "user id is empty",
			Error:      "Unauthorized Access",
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	data, err := uc.userService.UpdateLanguages(&languages, userId)
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

	c.JSON(http.StatusOK, data)
}
