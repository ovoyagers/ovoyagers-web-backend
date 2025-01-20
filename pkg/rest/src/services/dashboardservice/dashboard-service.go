package dashboardservice

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/dashboarddao"
)

type DashboardService struct {
	dashboardDao *dashboarddao.DashboardDao
}

func NewDashboardService(globalCfg *config.GlobalConfig) *DashboardService {
	return &DashboardService{dashboardDao: dashboarddao.NewDashboardDao(globalCfg)}
}
