package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/controllers/formcontroller"
)

type FormRoute struct {
	formController *formcontroller.FormController
}

func NewFormRoute(globalCfg *config.GlobalConfig) FormRoute {
	formController := formcontroller.NewFormController(globalCfg)
	return FormRoute{formController: formController}
}

func (ar FormRoute) SetupFormRoute(rg *gin.RouterGroup) {
	router := rg.Group("/form")

	router.POST("/create", ar.formController.CreateForm)
}