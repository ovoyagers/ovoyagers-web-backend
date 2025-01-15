package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/controllers/invitecontroller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type InviteRoute struct {
	invitecontroller *invitecontroller.InviteController
}

func NewInviteRoute(globalCfg *config.GlobalConfig) InviteRoute {
	ic := invitecontroller.NewInviteController(globalCfg)
	return InviteRoute{invitecontroller: ic}
}

func (ir *InviteRoute) SetupInviteRoute(rg *gin.RouterGroup) {
	router := rg.Group("/invite")
	router.Use(middlewares.DeserializeUser())

	router.POST("/send-invite", ir.invitecontroller.SendInvite)
}
