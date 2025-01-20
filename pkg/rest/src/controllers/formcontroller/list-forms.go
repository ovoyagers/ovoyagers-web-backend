package formcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// GetForms returns a list of forms in the database, with pagination support.
//
//	@Summary		Get forms
//	@Description	Get a list of forms in the database, with pagination support
//	@Tags			form
//	@Produce		json
//	@Param			limit	query		int	false	"Limit of the number of records to return"
//	@Param			page	query		int	false	"Page of the number of records to return"
//	@Success		200		{object}	models.Response
//	@Failure		422		{object}	models.Error
//	@Failure		409		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/form/all [get]
//	@Security		BearerAuth
func (fc *FormController) GetForms(c *gin.Context) {
	limit, page, offset, err := parsePaginationParams(c)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Invalid pagination parameters")
		return
	}

	// Call the service to get forms and total count
	forms, totalCount, err := fc.formSvc.GetForms(limit, page, offset)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to get forms")
		return
	}

	// Prepare response
	data := prepareResponse(forms, totalCount, limit, page)

	// Return response
	utils.HTTPResponseHandler(c, data, http.StatusOK, "Forms retrieved successfully")
}

// GetFormsByCategory returns a list of forms in the database, with pagination support, filtered by category.
//
//	@Summary		Get forms by category
//	@Description	Get a list of forms in the database, with pagination support, filtered by category
//	@Tags			form
//	@Produce		json
//	@Param			category	query		string	true	"Category of the forms"
//	@Param			limit		query		int		false	"Limit of the number of records to return"
//	@Param			page		query		int		false	"Page of the number of records to return"
//	@Success		200			{object}	models.Response
//	@Failure		422			{object}	models.Error
//	@Failure		409			{object}	models.Error
//	@Failure		400			{object}	models.Error
//	@Failure		500			{object}	models.Error
//	@Router			/form/by-category [get]
//	@Security		BearerAuth
func (fc *FormController) GetFormsByCategory(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		utils.HTTPErrorHandler(c, nil, http.StatusBadRequest, "Category is required")
		return
	}

	if err := validateCategory(category); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, err.Error())
		return
	}

	limit, page, offset, err := parsePaginationParams(c)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, err.Error())
		return
	}

	// Call the service to get forms and total count
	forms, totalCount, err := fc.formSvc.GetFormsByCategory(category, limit, page, offset)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to get forms")
		return
	}

	// Prepare response
	data := prepareResponse(forms, totalCount, limit, page)

	// Return response
	utils.HTTPResponseHandler(c, data, http.StatusOK, "Forms retrieved successfully")
}
