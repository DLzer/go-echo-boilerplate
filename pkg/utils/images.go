package utils

import (
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/DLzer/go-echo-boilerplate/pkg/httpErrors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var allowedImagesContentType = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
}

func determineFileContentType(fileHeader textproto.MIMEHeader) (string, error) {
	contentTypes := fileHeader["Content-Type"]
	if len(contentTypes) < 1 {
		return "", httpErrors.ErrNotAllowedImageHeader
	}
	return contentTypes[0], nil
}

func CheckImageContentType(image *multipart.FileHeader) error {
	// Check content type from header
	if !IsAllowedImageHeader(image) {
		return httpErrors.ErrNotAllowedImageHeader
	}

	// Check real content type
	imageFile, err := image.Open()
	if err != nil {
		return httpErrors.ErrBadRequest
	}
	defer imageFile.Close()

	fileHeader := make([]byte, 512)
	if _, err = imageFile.Read(fileHeader); err != nil {
		return httpErrors.ErrBadRequest
	}

	if !IsAllowedImageContentType(fileHeader) {
		return httpErrors.ErrNotAllowedImageHeader
	}
	return nil
}

func IsAllowedImageHeader(image *multipart.FileHeader) bool {
	contentType, err := determineFileContentType(image.Header)
	if err != nil {
		return false
	}
	_, allowed := allowedImagesContentType[contentType]
	return allowed
}

func GetImageExtension(image *multipart.FileHeader) (string, error) {
	contentType, err := determineFileContentType(image.Header)
	if err != nil {
		return "", err
	}

	extension, has := allowedImagesContentType[contentType]
	if !has {
		return "", errors.New("prohibited image extension")
	}
	return extension, nil
}

func GetImageContentType(image []byte) (string, bool) {
	contentType := http.DetectContentType(image)
	extension, allowed := allowedImagesContentType[contentType]
	return extension, allowed
}

func IsAllowedImageContentType(image []byte) bool {
	_, allowed := GetImageContentType(image)
	return allowed
}

func GetUniqFileName(userID string, fileExtension string) string {
	randString := uuid.New().String()
	return "userid_" + userID + "_" + randString + "." + fileExtension
}
