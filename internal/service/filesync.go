package service

import (
	"context"

	fs "github.com/underthetreee/fsync/pkg/proto"
)

type FileManager interface {
	CreateFile(file *fs.File) error
	GetFile(filename string) (*fs.File, error)
	DeleteFile(filename string) error
}

type Producer interface {
	ProduceFileEvent(ctx context.Context, topic string, event *fs.FileEvent) error
}

type FileSyncService struct {
	mng  FileManager
	prod Producer
}

func NewFileSyncService(manager FileManager, producer Producer) *FileSyncService {
	return &FileSyncService{
		mng:  manager,
		prod: producer,
	}
}

func (s *FileSyncService) UploadFile(ctx context.Context, file *fs.File) error {
	if err := s.mng.CreateFile(file); err != nil {
		return err
	}

	event := NewFileEvent(file.Filename, UPLOAD)
	if err := s.prod.ProduceFileEvent(ctx, "file-upload", event); err != nil {
		return err
	}
	return nil
}

func (s *FileSyncService) DownloadFile(ctx context.Context, filename string) (*fs.File, error) {
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
