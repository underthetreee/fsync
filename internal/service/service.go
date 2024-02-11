package service

import (
	"context"

	"github.com/underthetreee/fsync/internal/model"
	"github.com/underthetreee/fsync/pkg/proto"
)

type FileManager interface {
	CreateFile(file *model.File) error
}

type FileSyncService struct {
	mng FileManager
}

func NewFileSyncService(manager FileManager) *FileSyncService {
	return &FileSyncService{
		mng: manager,
	}
}

func (s *FileSyncService) UploadFile(ctx context.Context, protoFile *proto.File) error {
	file, err := model.ToModel(protoFile)
	if err != nil {
		return err
	}
	if err := s.mng.CreateFile(file); err != nil {
		return err
	}
	return nil
}
