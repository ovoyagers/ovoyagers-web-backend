package followservice

import (
	"context"

	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/followdao"
)

type FollowService struct {
	followDao *followdao.FollowDao
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{
		followDao: followdao.NewFolloWDao(ctx),
	}
}

func (fs *FollowService) CreateFollowRequest(followUsername string, userid string) (map[string]interface{}, error) {
	return fs.followDao.CreateFollowRequest(followUsername, userid)
}

func (fs *FollowService) CancelFollowRequest(followerUsername string, userid string) error {
	return fs.followDao.CancelFollowRequest(followerUsername, userid)
}

func (fs *FollowService) AcceptFollowRequest(followUsername string, userid string) (map[string]interface{}, error) {
	return fs.followDao.AcceptFollowRequest(followUsername, userid)
}
