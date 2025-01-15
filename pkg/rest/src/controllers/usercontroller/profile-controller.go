package usercontroller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/petmeds24/backend/pkg/rest/src/models"
	"github.com/petmeds24/backend/pkg/rest/src/models/usermodel"
	"github.com/petmeds24/backend/pkg/rest/src/utils"
)

// GetProfileInfo gets a user profile info
//
//	@Summary		get a user profile info
//	@Description	get a user profile info
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		422	{object}	models.Error
//	@Failure		400	{object}	models.Error
//	@Failure		400	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/user/me [get]
//	@Security		BearerAuth
func (uc *UserController) GetProfileInfo(ctx *gin.Context) {
	var userid = ctx.GetString("user_id")
	if userid == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
			Message:    "no token provided",
			Error:      "Status Unauthorized",
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	user, err := uc.userService.GetProfileInfo(userid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.Error{
			Message:    err.Error(),
			Error:      "Status Not Found",
			Status:     "error",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateProfileInfo(ctx *gin.Context) {
	var userProfile usermodel.UserProfile
	if err := ctx.ShouldBindJSON(&userProfile); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.Error{
			Message:    err.Error(),
			Error:      "Status Unprocessable Entity",
			Status:     "error",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}

	if err := userProfile.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Error{
			Message:    err.Error(),
			Error:      "Status Bad Request",
			Status:     "error",
			StatusCode: http.StatusBadRequest,
		})
	}

	var userid = ctx.GetString("user_id")
	if userid == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
			Message:    "no token provided",
			Error:      "Status Unauthorized",
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}

	editProfile := userProfile.ConvertStructToMap()
	// update user profile
	editProfile["id"] = userid

	profile, err := uc.userService.EditProfile(editProfile)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.Error{
			Message:    err.Error(),
			Error:      "Status Not Found",
			Status:     "error",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// UpdateProfilePicture updates a user profile picture
//
//	@Summary		updates a user profile picture
//	@Description	updates a user profile picture
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			avatar	formData	file	true	"avatar"
//	@Failure		422		{object}	models.Error
//	@Failure		409		{object}	models.Error
//	@Failure		400		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/user/upload-profile-picture [put]
//	@Security		BearerAuth
func (uc *UserController) UpdateProfilePicture(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusBadRequest, "Invalid file upload")
		return
	}

	userID, ok := ctx.Get("user_id")
	if !ok {
		utils.HTTPErrorHandler(ctx, errors.New("no token provided"), http.StatusUnauthorized, "Unauthorized access")
		return
	}

	// Check if the uploaded file is an image
	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		utils.HTTPErrorHandler(ctx, errors.New("uploaded file is not an image"), http.StatusBadRequest, "Invalid file type")
		return
	}

	// Check file size limit (10MB)
	if file.Size > 10*1024*1024 {
		utils.HTTPErrorHandler(ctx, errors.New("file size exceeds 10MB"), http.StatusBadRequest, "File too large")
		return
	}

	// Open file to read its content
	fileContent, err := file.Open()
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Failed to open uploaded file")
		return
	}
	defer fileContent.Close()

	// Read file content
	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Failed to read file content")
		return
	}

	// Detect MIME type from file content
	mimeType := http.DetectContentType(fileBytes)
	var base64Encoding string
	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	case "image/gif":
		base64Encoding = "data:image/gif;base64,"
	case "image/webp":
		base64Encoding = "data:image/webp;base64,"
	case "image/svg+xml":
		base64Encoding = "data:image/svg+xml;base64,"
	default:
		utils.HTTPErrorHandler(ctx, errors.New("unsupported image type"), http.StatusBadRequest, "Invalid image format")
		return
	}

	// Base64 encode the file content
	base64Encoding += base64.StdEncoding.EncodeToString(fileBytes)

	// Extract the file extension using path.Ext
	fileExt := path.Ext(file.Filename)

	// Prepare profile update payload
	updateProfile := map[string]string{
		"id":        userID.(string),
		"avatar":    base64Encoding,
		"avatar_id": file.Filename,
		"filename":  fmt.Sprintf("avatar%s", fileExt),
	}

	// Upload the profile picture
	profile, err := uc.userService.Upload(updateProfile)
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Failed to upload image")
		return
	}

	// Send success response
	utils.HTTPResponseHandler(ctx, profile, http.StatusOK, "Profile picture updated successfully")
}

// DeleteProfilePicture deletes a user profile picture
//
//	@Summary		deletes a user profile picture
//	@Description	deletes a user profile picture
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		500	{object}	models.Error
//	@Router			/user/delete-profile-picture [delete]
//	@Security		BearerAuth
func (uc *UserController) DeleteProfilePicture(ctx *gin.Context) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		utils.HTTPErrorHandler(ctx, errors.New("no token provided"), http.StatusUnauthorized, "Unauthorized access")
		return
	}

	data, err := uc.userService.DeleteAvatar(userID.(string))
	if err != nil {
		utils.HTTPErrorHandler(ctx, err, http.StatusInternalServerError, "Failed to delete avatar")
		return
	}

	utils.HTTPResponseHandler(ctx, data, http.StatusOK, "Avatar deleted successfully")
}
