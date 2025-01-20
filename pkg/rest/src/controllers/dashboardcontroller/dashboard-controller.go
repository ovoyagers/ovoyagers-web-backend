package dashboardcontroller

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/services/dashboardservice"
)

type DashboardController struct {
	dashboardSvc *dashboardservice.DashboardService
}

func NewDashboardController(globalCfg *config.GlobalConfig) *DashboardController {
	return &DashboardController{
		dashboardSvc: dashboardservice.NewDashboardService(globalCfg),
	}
}