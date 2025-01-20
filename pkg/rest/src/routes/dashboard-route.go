package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/controllers/dashboardcontroller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type DashboardRoute struct {
	dashboardController *dashboardcontroller.DashboardController
}

func NewDashboardRoute(globalCfg *config.GlobalConfig) DashboardRoute {
	dashboardController := dashboardcontroller.NewDashboardController(globalCfg)
	return DashboardRoute{
		dashboardController: dashboardController,
	}
}

func (dr DashboardRoute) SetupDashboardRoute(rg *gin.RouterGroup) {
	router := rg.Group("/dashboard")

	router.Use(middlewares.DeserializeUser())
	router.GET("/web-analytics", dr.dashboardController.GetWebAnalytics)
}
