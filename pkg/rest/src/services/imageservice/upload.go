package imageservice

import (
	"encoding/json"

	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	log "github.com/sirupsen/logrus"
)

func (s *ImageKit) Upload(filemap map[string]string) (map[string]interface{}, error) {
	user, err := s.userDao.GetUserById(filemap["id"])
	// if user is not found
	if err != nil {
		return nil, err
	}

	// convert string to authmodel.ProfilePicture
	ppic := user["profilePicture"].(string)
	if ppic == "" {
		return nil, nil
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
	if err != nil && res != nil {
		return nil, err
	}

	// upload new profile picture
	resp, err := s.imgKit.Uploader.Upload(ctx, filemap["avatar"], uploader.UploadParam{
		FileName: filemap["filename"],
		Folder:   "/" + filemap["id"],
	})

	if err != nil {
		return nil, err
	}
	data, err := s.userDao.UpdateProfilePicture(resp, filemap["id"])
	if err != nil {
		return nil, err
	}

	return data, nil
}
