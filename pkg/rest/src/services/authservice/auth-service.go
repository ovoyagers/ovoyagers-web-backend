package authservice

import (
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/daos/handlers/authdao"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	"github.com/petmeds24/backend/pkg/rest/src/services/cacheservice"
)

type AuthService struct {
	authDao *authdao.AuthDao
	cache   *cacheservice.CacheService
}

func NewAuthService(globalCfg *config.GlobalConfig) *AuthService {
	return &AuthService{
		authDao: authdao.NewAuthDao(globalCfg),
		cache:   cacheservice.NewCacheService(globalCfg),
	}
}

func (as *AuthService) RegisterUser(user *authmodel.RegisterRequest, otp string) (map[string]interface{}, error) {
	return as.authDao.RegisterUser(user, otp)
}

func (as *AuthService) GetUserByPhone(phone string) (map[string]interface{}, int, error) {
	return as.authDao.GetUserByPhone(phone)
}

func (as *AuthService) GetUserByEmail(email string) (map[string]interface{}, error) {
	return as.authDao.GetUserByEmail(email)
}

func (as *AuthService) IsUserExists(user authmodel.CheckUser) (bool, error) {
	return as.authDao.IsUserExists(user)
}
func (as *AuthService) UpdateVerifiedUser(petid string) (map[string]interface{}, error) {
	return as.authDao.UpdateVerifiedUser(petid)
}

func (as *AuthService) FindUserByPhone(phone string) (map[string]interface{}, error) {
	return as.authDao.FindUserByPhone(phone)
}

func (as *AuthService) FindUserByID(id string) (map[string]interface{}, error) {
	return as.authDao.FindUserByID(id)
}

func (as *AuthService) InsertUser(user authmodel.User) (map[string]interface{}, error) {
	return as.authDao.InsertUser(user)
}

func (as *AuthService) GenerateRandomPetName(petname string) (string, int, error) {
	return "", 0, nil
}

func (as *AuthService) ForgetPassword(email string) (map[string]interface{}, error) {
	return as.authDao.ForgetPassword(email)
}
