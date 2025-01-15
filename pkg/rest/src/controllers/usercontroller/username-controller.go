package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	log "github.com/sirupsen/logrus"
)

// GenerateRandomUsername godoc
//
//	@Summary		Generate Random Username
//	@Description	Generate Random Username
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			fullname	query		usermodel.Name	true	"Full Name"
//	@Success		200			{object}	models.Response
//	@Failure		422			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/user/get-random-username [get]
//	@Security		BearerAuth
func (uc *UserController) GenerateRandomUsername(c *gin.Context) {
	// fetch user data from request body
	var name usermodel.Name
	name.FullName = c.Query("fullname")

	// validate user data
	err := name.Validate()
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

	// generate random username
	randUsernames := uc.userService.GetRandomUsernames(name.FullName)
	c.JSON(http.StatusOK, models.Response{
		Data:       randUsernames,
		Message:    "Random usernames generated successfully",
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}

// UpdateUsername godoc
//
//	@Summary		Update Username
//	@Description	Update Username
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			fullname	body		usermodel.Name	true	"Full Name"
//	@Success		200			{object}	models.Response
//	@Failure		422			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/user/update-username [put]
//	@Security		BearerAuth
func (uc *UserController) UpdateUsername(c *gin.Context) {
	// fetch user data from request body
	var name usermodel.Name
	err := c.ShouldBindJSON(&name)
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
	// validate user data
	err = name.Validate()
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

	userid := c.GetString("user_id")
	user, err := uc.userService.UpdateUsername(name.FullName, userid)
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
		Data:       user,
		Message:    "Username updated successfully",
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}
