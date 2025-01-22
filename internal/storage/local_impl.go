package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

// NewLocalStorage creates a new LocalStorage instance.
func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{BasePath: basePath}
}

// UploadFile saves the file to the local filesystem.
func (localStorage *LocalStorage) UploadFile(file multipart.File, fileName string) (string, error) {
	defer file.Close()
	path := filepath.Join(localStorage.BasePath, fileName)
	outFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return path, nil
}

// DeleteFile removes the file from the local filesystem.
func (localStorage *LocalStorage) DeleteFile(fileName string) error {
	path := filepath.Join(localStorage.BasePath, fileName)
	return os.Remove(path)
}

// GetFileURL returns the local file path.
func (localStorage *LocalStorage) GetFileURL(fileName string) (string, error) {
	return filepath.Join(localStorage.BasePath, fileName), nil
}
