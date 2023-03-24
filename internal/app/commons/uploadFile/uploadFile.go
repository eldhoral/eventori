package uploadFile

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	OtherDocument = iota
)

var (
	AllowedFile  = []string{"jpg", "png", "jpeg", "svg"}
	DocumentPath = map[int]string{
		OtherDocument: "./assets/upload/image/",
	}
)

func UploadFIleInstance(fileHeader *multipart.FileHeader, c *gin.Context) (destinationFile string, err error) {
	filename := fileHeader.Filename
	extension := filepath.Ext(filename)
	ext := strings.ReplaceAll(filepath.Ext(filename), ".", "")

	if !InArray(ext, AllowedFile) {
		return "", errors.New("File extension must be between the following types: " + strings.Join(AllowedFile, ", "))
	}

	// Avoid path traversal
	fileRe := regexp.MustCompile(`^\.*/|/|\.\..*$`)
	match := fileRe.MatchString(filename)
	if match {
		return "", fmt.Errorf("File name is invalid. Please use valid character: %s", filename)
	}

	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension
	destinationFile = DocumentPath[OtherDocument] + newFileName

	//create file and copy
	err = c.SaveUploadedFile(fileHeader, destinationFile)
	if err != nil {
		return "", fmt.Errorf("Failed saving image to destinationFile: %s", filename)
	}

	return
}

// This function is the same as in_array in PHP
func InArray(key interface{}, array interface{}) bool {
	switch key := key.(type) {
	case string:
		for _, item := range array.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range array.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range array.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}
