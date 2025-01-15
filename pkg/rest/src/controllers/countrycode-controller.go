package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/utils/constants"
	log "github.com/sirupsen/logrus"
)

type CountryCodeController struct {
}

func NewCountryCodeController() *CountryCodeController {
	return &CountryCodeController{}
}

// GetCountryCodes returns all country codes
//
//	@Summary		Get all country codes
//	@Description	Get all country codes
//	@Tags			Country Codes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/countries/all [get]
func (cc *CountryCodeController) GetCountryCodes(c *gin.Context) {
	countryCodes, err := constants.CountryData()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Error{
			Message:    err.Error(),
			Error:      "Internal Server Error",
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message:    "country codes retrieved successfully",
		Data:       countryCodes,
		Status:     "success",
		StatusCode: http.StatusOK,
	})
}
