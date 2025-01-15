package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// Helper function to handle passwordless login
func handlePasswordlessLogin(c *gin.Context, ac *AuthController, email string) error {
	code, err := utils.GenerateOTP(6)
	if err != nil {
		return err
	}
	if err = ac.authService.UpdateOTP(email, code); err != nil {
		return err
	}
	if err = ac.authService.SendOTPViaEmail(email, code); err != nil {
		return err
	}
	c.JSON(http.StatusOK, models.Response{
		Message:    "OTP sent successfully",
		Data:       map[string]interface{}{"code": code},
		Status:     "success",
		StatusCode: http.StatusOK,
	})
	return nil
}
