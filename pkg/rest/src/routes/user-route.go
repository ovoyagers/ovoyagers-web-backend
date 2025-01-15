package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/controllers/usercontroller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type UserRoute struct {
	userController *usercontroller.UserController
}

func NewUserRoute(globalCfg *config.GlobalConfig) UserRoute {
	uc := usercontroller.NewUserController(globalCfg)
	return UserRoute{userController: uc}
}

func (ur UserRoute) SetupUserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/user")
	router.Use(middlewares.DeserializeUser())

	router.PUT("/update-about", ur.userController.UpdateAboutUser)
	router.GET("/me", ur.userController.GetProfileInfo)
	router.PUT("/update-user", ur.userController.UpdateUser)
	router.PUT("/update-languages", ur.userController.UpdateLanguages)
	router.GET("/get-random-username", ur.userController.GenerateRandomUsername)
	router.PUT("/update-username", ur.userController.UpdateUsername)
	router.PUT("/upload-profile-picture", ur.userController.UpdateProfilePicture)
	router.DELETE("/delete-profile-picture", ur.userController.DeleteProfilePicture)
}
