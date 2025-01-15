package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	log "github.com/sirupsen/logrus"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrNoDataFound   = errors.New("no data found")
)

func HTTPErrorHandler(c *gin.Context, err error, statusCode int, message string) {
	log.Error(err)
	c.JSON(statusCode, models.Error{
		Message:    err.Error(),
		Error:      message,
		Status:     "error",
		StatusCode: statusCode,
	})
}

func HTTPResponseHandler(c *gin.Context, data interface{}, statusCode int, message string) {
	c.JSON(statusCode, models.Response{
		Message:    message,
		Data:       data,
		Status:     "success",
		StatusCode: statusCode,
	})
}

func HTTPErrorWithDataHandler(c *gin.Context, statusCode int, message string, data interface{}) {
	log.Error(data)
	c.JSON(statusCode, models.Error{
		Data:       data,
		Message:    message,
		Status:     "error",
		StatusCode: statusCode,
	})
}
