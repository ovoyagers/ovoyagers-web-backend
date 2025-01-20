package authcontroller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/config"
	"github.com/petmeds24/backend/pkg/rest/src/models/authmodel"
	"github.com/petmeds24/backend/pkg/rest/src/services/authservice"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
	"github.com/petmeds24/backend/pkg/rest/src/utils/constants"
	log "github.com/sirupsen/logrus"
)

type AuthController struct {
	authService *authservice.AuthService
	consts      *constants.Constants
}

var (
	ACCESS_TOKEN_EXPIRY  = 15 * time.Minute    // 15 minutes
	REFRESH_TOKEN_EXPIRY = 15 * time.Hour * 24 // 15 days
)

func NewAuthController(globalCfg *config.GlobalConfig) *AuthController {
	return &AuthController{
		authService: authservice.NewAuthService(globalCfg),
		consts:      constants.GetConstants(globalCfg.GetConfig()),
	}
}

// RegisterUser creates a new User
//
//	@Summary		Create a new User
//	@Description	Create a new User
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		authmodel.RegisterRequest	true	"User"
//	@Failure		422		{object}	models.Error
//	@Failure		409		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/auth/create [post]
func (ac *AuthController) RegisterUser(c *gin.Context) {
	// Fetch user data from request body
	var user authmodel.RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusUnprocessableEntity, "Invalid input data")
		return
	}

	// Validate user data
	if err := user.Validate(); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Validation failed")
		return
	}

	// Check if the user already exists
	userExists, err := ac.authService.GetUserByEmail(user.Email)
	if err == nil && userExists != nil {
		message := "User already exists"
		if !userExists["isVerified"].(bool) {
			message += " and is not verified"
		} else {
			message += " and is verified"
		}
		utils.HTTPErrorWithDataHandler(c, http.StatusConflict, message, userExists)
		return
	}
	// Generate OTP
	code, err := utils.GenerateOTP(6)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to generate OTP")
		return
	}

	// Create a channel for error handling
	errCh := make(chan error, 2)

	// Parallelize user creation and OTP email sending
	var userCreated interface{}
	go func() {
		createdUser, err := ac.authService.RegisterUser(&user, code)
		if err != nil {
			errCh <- fmt.Errorf("user_creation:%w", err)
			return
		}
		userCreated = createdUser
		errCh <- nil
	}()

	go func() {
		err := ac.authService.SendOTPViaEmail(user.Email, code)
		if err != nil {
			errCh <- fmt.Errorf("email_sending:%w", err)
			return
		}
		errCh <- nil
	}()

	// Collect errors from both goroutines
	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			if strings.Contains(err.Error(), "user_creation") {
				utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to create user in the database")
				return
			}
			if strings.Contains(err.Error(), "email_sending") {
				utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to send OTP via email")
				return
			}
		}
	}

	// If all succeeded, return the success response
	utils.HTTPResponseHandler(c, userCreated, http.StatusCreated, "User created successfully")
}

// VerifyEmail verifies an email
//
//	@Summary		Verify an email
//	@Description	Verify an email
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			verify	body		authmodel.VerifyEmailRequest	true	"Email"
//	@Failure		422		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/auth/verify-email [post]
func (ac *AuthController) VerifyEmail(c *gin.Context) {
	var verify authmodel.VerifyEmailRequest
	if err := c.ShouldBindJSON(&verify); err != nil {
		log.Error(err)
		utils.HTTPErrorHandler(c, err, http.StatusUnprocessableEntity, "Invalid input data")
		return
	}

	user, err := ac.authService.VerifyOTPViaEmail(verify)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Invalid OTP")
		return
	}

	// set user to cache
	err = ac.authService.SetUserToCache(user)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to set user to cache")
		return
	}

	// generate token
	jwtUtil := utils.NewJWTUtil()
	token, err := jwtUtil.CreateToken(user["id"].(string), user["email"].(string))
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to generate token")
		return
	}
	c.SetCookie("access_token", token.AccessToken, int(ACCESS_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("refresh_token", token.RefreshToken, int(REFRESH_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("logged_in", "true", int(ACCESS_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	utils.HTTPResponseHandler(c, user, http.StatusOK, "Email verified successfully")
}

// ResendEmailOTP resends an email OTP
//
//	@Summary		Resend Email OTP
//	@Description	Resend Email OTP
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			email	body		authmodel.ResendEmailOTP	true	"Email"
//	@Failure		422		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/auth/resend-email-otp [post]
func (ac *AuthController) ResendEmailOTP(c *gin.Context) {
	var resend authmodel.ResendEmailOTP
	if err := c.ShouldBindJSON(&resend); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusUnprocessableEntity, "Invalid input data")
		return
	}

	err := resend.Validate()
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Validation failed")
		return
	}

	// Check if user exists
	user, err := ac.authService.CheckUser(resend.Email)
	if err != nil || user == nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "User not found")
		return
	}

	// Generate OTP
	code, err := utils.GenerateOTP(6)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Failed to generate OTP")
		return
	}

	// Update OTP asynchronously
	go func(email string, otpCode string) {
		// Update OTP in database
		if err := ac.authService.UpdateOTP(email, otpCode); err != nil {
			// Log the error for debugging purposes
			fmt.Printf("Error updating OTP: %v\n", err)
			return
		}

		// Send OTP via email
		if err := ac.authService.SendOTPViaEmail(email, otpCode); err != nil {
			// Log the error for debugging purposes
			fmt.Printf("Error sending OTP email: %v\n", err)
			return
		}
	}(resend.Email, code)

	// Respond immediately to the client
	data := map[string]interface{}{
		"email":      resend.Email,
		"id":         user["id"],
		"isVerified": user["isVerified"],
	}

	utils.HTTPResponseHandler(c, data, http.StatusOK, "OTP sent successfully")
}

// LoginEmail sign in a the existing user
//
//	@Summary		Login a user
//	@Description	Login a user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		authmodel.LoginEmailRequest	true	"User"
//	@Failure		422		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/auth/login [post]
func (ac *AuthController) LoginUser(c *gin.Context) {
	var login authmodel.LoginEmailRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusUnprocessableEntity, "Status Unprocessable Entity")
		return
	}

	if !login.Passwordless && login.Password == "" {
		utils.HTTPErrorHandler(c, errors.New("password is required"), http.StatusBadRequest, "Status Bad Request")
		return
	}

	// validate user data
	if err := login.Validate(); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Status Bad Request")
		return
	}

	// check user exists
	user, err := ac.authService.CheckUser(login.Email)
	if err != nil || len(user) < 1 {
		utils.HTTPErrorHandler(c, err, http.StatusNotFound, "Status Not Found")
		return
	}

	if !user["isVerified"].(bool) {
		utils.HTTPErrorHandler(c, errors.New("user is not verified"), http.StatusUnauthorized, "Status Unauthorized")
		return
	}

	if login.Passwordless {
		if err := handlePasswordlessLogin(c, ac, login.Email); err != nil {
			utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		return
	}

	// check password
	if !utils.VerifyPassword(login.Password, user["password"].(string)) {
		utils.HTTPErrorHandler(c, errors.New("invalid password"), http.StatusUnauthorized, "Status Unauthorized")
		return
	}

	// generate token
	jwtUtil := utils.NewJWTUtil()
	token, err := jwtUtil.CreateToken(user["id"].(string), user["email"].(string))
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	c.SetCookie("access_token", token.AccessToken, int(ACCESS_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("refresh_token", token.RefreshToken, int(REFRESH_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("logged_in", "true", int(ACCESS_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	utils.HTTPResponseHandler(c, data, http.StatusOK, "Login successful")
}

// Refresh is a method to refresh the access token and refresh token
//
//	@Summary		Refresh
//	@Description	Refresh
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		422	{object}	models.Error
//	@Failure		403	{object}	models.Error
//	@Failure		400	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/auth/refresh [get]
func (ac *AuthController) RefreshTokens(c *gin.Context) {
	jwtUtil := utils.NewJWTUtil()
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		if refreshToken == "" {
			refreshToken = c.Request.Header.Get("x-refresh-token")
			if refreshToken == "" {
				utils.HTTPErrorHandler(c, err, http.StatusForbidden, "refresh token not found")
				return
			}
		}
	}

	claims, err := jwtUtil.ValidateRefreshToken(refreshToken)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusForbidden, "refresh token is invalid")
		return
	}

	user, err := ac.authService.CheckUser(claims["email"].(string))
	if err != nil || len(user) < 1 {
		utils.HTTPErrorHandler(c, err, http.StatusNotFound, "user not found or incorrect")
		return
	}

	if len(user) == 0 {
		utils.HTTPErrorHandler(c, fmt.Errorf("the user belonging to this token no longer exists"), http.StatusForbidden, "Status forbidden")
		return
	}

	newToken, err := jwtUtil.CreateToken(user["id"].(string), user["email"].(string))
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	c.SetCookie("access_token", newToken.AccessToken, int(ACCESS_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("refresh_token", newToken.RefreshToken, int(REFRESH_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("logged_in", "true", int(ACCESS_TOKEN_EXPIRY.Seconds()), "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	utils.HTTPResponseHandler(c, newToken, http.StatusOK, "Tokens refreshed successfully")
}

// Logout is a method to logout a user
//
//	@Summary		Logout
//	@Description	Logout
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		422	{object}	models.Error
//	@Failure		403	{object}	models.Error
//	@Failure		400	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/auth/logout [get]
//	@Security		BearerAuth
func (ac *AuthController) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("refresh_token", "", -1, "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("logged_in", "", -1, "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	utils.HTTPResponseHandler(c, nil, http.StatusOK, "Logout successful")
}

// ForgotPassword is a method to send a reset password link to the user's email
//
//	@Summary		ForgotPassword
//	@Description	ForgotPassword
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			email	body		authmodel.ForgetPasswordRequest	true	"Email"
//	@Success		200		{object}	models.Response
//	@Failure		422		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/auth/forgot-password [post]
//	@Security		BearerAuth
func (ac *AuthController) ForgetPassword(c *gin.Context) {
	var forgetPassword authmodel.ForgetPasswordRequest
	if err := c.ShouldBindJSON(&forgetPassword); err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Bad Request")
		return
	}
	err := forgetPassword.Validate()
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusBadRequest, "Bad Request")
		return
	}
	data, err := ac.authService.ForgetPassword(forgetPassword.Email)
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// generate token
	jwtUtil := utils.NewJWTUtil()
	token, err := jwtUtil.CreateToken(data["id"].(string), data["email"].(string))
	if err != nil {
		utils.HTTPErrorHandler(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.SetCookie("access_token", token.AccessToken, 900, "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("refresh_token", token.RefreshToken, 900, "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	c.SetCookie("logged_in", "true", 900, "/", ac.consts.HOST, ac.consts.IS_SECURE, true)
	utils.HTTPResponseHandler(c, data, http.StatusOK, "Reset password link sent successfully")
}
