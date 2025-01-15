package authservice

import (
	"github.com/petmeds24/backend/pkg/rest/src/utils"
	"github.com/petmeds24/backend/pkg/rest/src/utils/constants"
)

func (as *AuthService) CheckUser(email string) (map[string]interface{}, error) {
	var (
		user map[string]interface{}
		err  error
	)
	// check if user exists in cache
	user, err = as.cache.GetRedisKey(email)
	// if user exists in cache, return user
	if err == nil {
		return user, nil
	}
	// if user does not exist in cache, check if user exists in database
	user, err = as.authDao.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// convert user to json
	userJson, err := utils.MapToJson(user)
	if err != nil {
		return nil, err
	}
	// if user exists in database, set user in cache
	err = as.cache.SetRedisKey(email, userJson, constants.Day)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (as *AuthService) SetUserToCache(user map[string]interface{}) error {
	// convert user to json
	userJson, err := utils.MapToJson(user)
	if err != nil {
		return err
	}
	// set user in cache
	err = as.cache.SetRedisKey(user["email"].(string), userJson, constants.Day)
	if err != nil {
		return err
	}
	return nil
}
