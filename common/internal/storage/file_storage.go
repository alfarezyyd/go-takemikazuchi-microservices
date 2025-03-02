package storage

import (
	"fmt"
	"github.com/spf13/viper"
	"mime/multipart"
	"os"
)

// FileStorage defines the methods for file storage operations.
type FileStorage interface {
	UploadFile(file multipart.File, fileName string) (string, error) // Upload file and return URL
	DeleteFile(fileName string) error                                // Delete file
	GetFileURL(fileName string) (string, error)                      // Get file URL
}

func ProvideFileStorage(viperConfig *viper.Viper) FileStorage {
	storageType := viperConfig.GetString("STORAGE_BACKEND")
	switch storageType {
	default:
		workingDirectory, _ := os.Getwd()
		return NewLocalStorage(fmt.Sprintf("%s/%s", workingDirectory, viperConfig.GetString("BASE_PATH")))
	}
}
