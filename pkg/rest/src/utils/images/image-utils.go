package utils

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/petmeds24/backend/pkg/rest/src/models"
)

func ConvertImageToBase64(file *multipart.FileHeader, filename string) (*models.ImgMetaData, error) {
	// Open file to read its content
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	// Read file content
	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return nil, err
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
		return nil, errors.New("invalid file type")
	}

	// Base64 encode the file content
	base64Encoding += base64.StdEncoding.EncodeToString(fileBytes)

	return &models.ImgMetaData{
		UserId:   "",
		Avatar:   base64Encoding,
		ImageId:  file.Filename,
		Filename: file.Filename,
	}, nil
}
