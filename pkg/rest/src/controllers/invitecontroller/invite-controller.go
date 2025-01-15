package invitecontroller

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/services/inviteservice"
)

type InviteController struct {
	inviteSvc *inviteservice.InviteService
}

func NewInviteController(globalCfg *config.GlobalConfig) *InviteController {
	return &InviteController{
		inviteSvc: inviteservice.NewInviteService(globalCfg),
	}
}
