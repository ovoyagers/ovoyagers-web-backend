package inviteservice

import (
	"github.com/petmeds24/backend/config"
	invitedao "github.com/petmeds24/backend/pkg/rest/src/daos/handlers/invite-dao"
	"github.com/petmeds24/backend/pkg/rest/src/services/imageservice"
)

type InviteService struct {
	imageKit  *imageservice.ImgKit
	inviteDao *invitedao.InviteDao
}

func NewInviteService(globalCfg *config.GlobalConfig) *InviteService {
	imgKit := imageservice.NewImgKit(globalCfg)
	return &InviteService{
		imageKit:  imgKit,
		inviteDao: invitedao.NewInviteDao(globalCfg),
	}
}
