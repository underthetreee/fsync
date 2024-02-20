package manager

import (
	"os"
	"path/filepath"

	"github.com/underthetreee/fsync/internal/model"
)

const (
	dirName         = ".fsync"
	filePermissions = 0644
	dirPermissions  = 0755
)

type Manager struct {
	StoragePath string
}

func NewManager() (*Manager, error) {
	storagePath, err := initStorage()
	if err != nil {
		return nil, err
	}
	return &Manager{
		StoragePath: storagePath,
	}, nil
}

func (m *Manager) GetFile(filename string) (*model.File, error) {
	filePath := filepath.Join(m.StoragePath, filename)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	file := &model.File{
		Filename: filename,
		Content:  fileContent,
	}
	return file, nil
}

func (m *Manager) StoreFile(file *model.File) error {
	filePath := filepath.Join(m.StoragePath, file.Filename)
	if err := os.WriteFile(filePath, file.Content, filePermissions); err != nil {
		return err
	}
	return nil
}

func (m *Manager) DeleteFile(filename string) error {
	filePath := filepath.Join(m.StoragePath, filename)
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func initStorage() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	storagePath := filepath.Join(homeDir, dirName)

	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		if err := os.Mkdir(storagePath, dirPermissions); err != nil {
			return "", err
		}
	}
	return storagePath, nil
}
