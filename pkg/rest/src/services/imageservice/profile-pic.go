package imageservice

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

func (s *ImageKit) DeleteAvatar(userid string) (map[string]interface{}, error) {
	user, err := s.userDao.GetUserById(userid)
	// if user is not found
	if err != nil {
		return nil, err
	}
	// convert string to authmodel.ProfilePicture
	ppic := user["profilePicture"].(string)
	if ppic == "" {
		return nil, errors.New("no profile picture found")
	}
	// convert string to authmodel.ProfilePicture
	var profile authmodel.ProfilePicture
	err = json.Unmarshal([]byte(ppic), &profile)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// delete profile picture
	res, err := s.imgKit.Media.DeleteFile(ctx, profile.FileId)
	// if file is not found
	log.Error(res)
	if err != nil {
		return nil, err
	}
	data, err := s.userDao.DeleteProfilePicture(userid)
	if err != nil {
		return nil, err
	}
	return data, nil
}
