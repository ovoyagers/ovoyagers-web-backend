package userservice

func (us *UserService) GetProfileInfo(userid string) (map[string]interface{}, error) {
	return us.userDao.GetProfileInfo(userid)
}

func (us *UserService) EditProfile(profile map[string]interface{}) (map[string]interface{}, error) {
	return us.userDao.EditProfile(profile)
}

func (us *UserService) Upload(filemap map[string]string) (map[string]interface{}, error) {
	return us.imgSrv.Upload(filemap)
}

func (us *UserService) DeleteAvatar(userid string) (map[string]interface{}, error) {
	return us.imgSrv.DeleteAvatar(userid)
}
