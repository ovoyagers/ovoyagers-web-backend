package userservice

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/userdao"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	"github.com/petmeds24/backend/pkg/rest/src/services/imageservice"
)

type UserService struct {
	userDao *userdao.UserDao
	imgSrv  *imageservice.ImageKit
}

func NewAuthService(globalCfg *config.GlobalConfig) *UserService {
	return &UserService{
		userDao: userdao.NewAuthDao(globalCfg),
		imgSrv:  imageservice.NewImageKit(globalCfg),
	}
}
func (us *UserService) UpdateAboutUser(aboutUser *usermodel.AboutUser, userid string) (map[string]interface{}, error) {
	return us.userDao.UpdateAboutUser(aboutUser, userid)
}

func (us *UserService) UpdateLanguages(lang *usermodel.Languages, userid string) (map[string]interface{}, error) {
	return us.userDao.UpdateLanguages(lang, userid)
}
