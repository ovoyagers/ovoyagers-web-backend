package inviteservice

import "github.com/petmeds24/backend/pkg/rest/src/models/invitemodel"

func (is *InviteService) SendInvite(invite invitemodel.InviteUser, userId string) error {
	return is.inviteDao.SendInvite(invite, userId)
}
