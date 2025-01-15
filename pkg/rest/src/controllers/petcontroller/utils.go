package petcontroller

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func (pc *PetController) convertImageToBase64(petImage *multipart.FileHeader) (string, error) {
	contentType := petImage.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return "", errors.New("invalid file type")
	}

	if petImage.Size > 5*1024*1024 {
		return "", errors.New("file size exceeds 5MB")
	}

	// open the file to read its content
	file, err := petImage.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read the file content into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(fileBytes)
	var base64Encoding string
	// determine the content type of the file
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	case "image/svg+xml":
		base64Encoding += "data:image/svg+xml;base64,"
	default:
		return "", errors.New("invalid file type")
	}

	base64Encoding += base64.StdEncoding.EncodeToString(fileBytes)

	return base64Encoding, nil
}
