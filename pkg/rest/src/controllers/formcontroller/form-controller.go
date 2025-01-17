package formcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/models/formmodel"
	"github.com/petmeds24/backend/pkg/rest/src/services/formservice"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

type FormController struct {
	formSvc *formservice.FormService
}

func NewFormController(globalCfg *config.GlobalConfig) *FormController {
	return &FormController{
		formSvc: formservice.NewFormService(globalCfg),
	}
}

// CreateForm creates a new form for ovoyagers website
//
//	@Summary		Create a new form
//	@Description	Create a new form for ovoyagers website with the specific category
//	@Tags			form
//	@Accept			json
//	@Produce		json
//	@Param			form	body		formmodel.Form	true	"Form"
//	@Failure		422		{object}	models.Error
//	@Failure		409		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/form/create [post]
func (fc *FormController) CreateForm(c *gin.Context) {
	var form formmodel.Form

	// Bind the JSON data from the request body to the form struct
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusUnprocessableEntity, "Invalid input data")
	}

	// Validate the form data
	if err := form.Validate(); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Invalid form data")
		return
	}

	// Create the form in the database
	if _, err := fc.formSvc.CreateForm(&form); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to create form")
		return
	}
	// Return a success response
	utils.HTTPResponseHandler(c, nil, http.StatusCreated, "Form created successfully")
}
