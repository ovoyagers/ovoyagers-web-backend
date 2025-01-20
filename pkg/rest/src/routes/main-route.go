package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
)

type MainRoute struct {
	globalCfg   *config.GlobalConfig
	routerGroup *gin.RouterGroup
}

func NewMainRoute(globalCfg *config.GlobalConfig, rg *gin.RouterGroup) MainRoute {
	return MainRoute{
		globalCfg:   globalCfg,
		routerGroup: rg,
	}
}

func (r MainRoute) SetupRoutes() {
	authRoute := NewAuthRoute(r.globalCfg)
	formRoute := NewFormRoute(r.globalCfg)
	dashboardRoute := NewDashboardRoute(r.globalCfg)
	userRoute := NewUserRoute(r.globalCfg)
	petRoute := NewPetRoute(r.globalCfg)
	recordRoute := NewRecordRoute(r.globalCfg)
	healthRoute := NewHealthRoute()
	countrycodeRoute := NewCountryCodeRoute()
	notificationRoute := NewNotificationRoute(r.globalCfg)
	inviteRoute := NewInviteRoute(r.globalCfg)

	healthRoute.SetupHealthRoute(r.routerGroup)
	countrycodeRoute.SetupCountryCodeRoute(r.routerGroup)
	authRoute.SetupAuthRoute(r.routerGroup)
	formRoute.SetupFormRoute(r.routerGroup)
	dashboardRoute.SetupDashboardRoute(r.routerGroup)
	userRoute.SetupUserRoute(r.routerGroup)
	petRoute.SetupPetRoute(r.routerGroup)
	recordRoute.SetupRecordRoute(r.routerGroup)
	notificationRoute.SetupNotificationRoute(r.routerGroup)
	inviteRoute.SetupInviteRoute(r.routerGroup)
}
