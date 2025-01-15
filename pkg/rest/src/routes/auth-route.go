package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/controllers/authcontroller"
	"github.com/petmeds24/backend/pkg/rest/src/middlewares"
)

type AuthRoute struct {
	authController *authcontroller.AuthController
}

func NewAuthRoute(globalCfg *config.GlobalConfig) AuthRoute {
	authController := authcontroller.NewAuthController(globalCfg)
	return AuthRoute{authController: authController}
}

func (ar AuthRoute) SetupAuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/create", ar.authController.RegisterUser)
	router.POST("/login", ar.authController.LoginUser)
	router.POST("/forgot-password", ar.authController.ForgetPassword)
	router.POST(("/verify-email"), ar.authController.VerifyEmail)
	router.POST("/resend-email-otp", ar.authController.ResendEmailOTP)
	router.GET("/refresh", ar.authController.RefreshTokens)
	router.GET("/logout", middlewares.DeserializeUser(), ar.authController.Logout)
}
