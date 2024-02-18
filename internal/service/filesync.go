package service

import (
	"context"

	"github.com/underthetreee/fsync/internal/model"
)

type FileManager interface {
	CreateFile(file *model.File) error
	GetFile(filename string) (*model.File, error)
	DeleteFile(filename string) error
}

type FileSyncService struct {
	mng FileManager
}

func NewFileSyncService(manager FileManager) *FileSyncService {
	return &FileSyncService{
		mng: manager,
	}
}

func (s *FileSyncService) UploadFile(ctx context.Context, file *model.File) error {
	if err := s.mng.CreateFile(file); err != nil {
		return err
	}
	return nil
}

func (s *FileSyncService) DownloadFile(ctx context.Context, filename string) (*model.File, error) {
	file, err := s.mng.GetFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *FileSyncService) DeleteFile(ctx context.Context, filename string) error {
	if err := s.mng.DeleteFile(filename); err != nil {
		return err
	}
	return nil
}
