package userservice

import "github.com/petmeds24/backend/pkg/rest/src/models/usermodel"

func (us *UserService) UpdateUser(user usermodel.UpdateUser, id string) (map[string]interface{}, error) {
	return us.userDao.UpdateUser(user, id)
}
