package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/controllers"
)

type CountryCodeRoute struct {
	countryCodeController *controllers.CountryCodeController
}

func NewCountryCodeRoute() *CountryCodeRoute {
	countryCodeController := controllers.NewCountryCodeController()
	return &CountryCodeRoute{countryCodeController: countryCodeController}
}

func (ccr CountryCodeRoute) SetupCountryCodeRoute(rg *gin.RouterGroup) {
	router := rg.Group("/countries")

	router.GET("/all", ccr.countryCodeController.GetCountryCodes)
}
