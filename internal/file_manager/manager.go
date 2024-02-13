package manager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/underthetreee/fsync/internal/model"
)

const dirName = ".fsync"

type Manager struct {
	StoragePath string
}

func NewManager() (*Manager, error) {
	storagePath, err := initStorage()
	if err != nil {
		return nil, err
	}
	log.Println("init file manager")

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

func (m *Manager) CreateFile(file *model.File) error {
	filePath := filepath.Join(m.StoragePath, file.Filename)
	if err := os.WriteFile(filePath, file.Content, 0644); err != nil {
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
		return "", nil
	}

	storagePath := filepath.Join(homeDir, dirName)

	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		if err := os.Mkdir(storagePath, 0755); err != nil {
			return "", fmt.Errorf("create directory: %w", err)
		}
	}
	return storagePath, nil
}
