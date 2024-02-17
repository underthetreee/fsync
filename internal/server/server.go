package server

import (
	"context"
	"log"

	"github.com/underthetreee/fsync/internal/model"
	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileSyncService interface {
	UploadFile(ctx context.Context, filename *model.File) error
	DownloadFile(ctx context.Context, filename string) (*model.File, error)
	DeleteFile(ctx context.Context, filename string) error
}

type Server struct {
	fs.UnimplementedFileSyncServiceServer
	svc FileSyncService
}

func Register(gRPCServer *grpc.Server, service FileSyncService) {
	fs.RegisterFileSyncServiceServer(gRPCServer, &Server{
		svc: service,
	})
}

func (s *Server) UploadFile(ctx context.Context, req *fs.UploadFileRequest,
) (*fs.UploadFileResponse, error) {
	file := model.ToModelFile(req.File)

	if err := s.svc.UploadFile(ctx, file); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &fs.UploadFileResponse{}, nil
}

func (s *Server) DownloadFile(ctx context.Context, req *fs.DownloadFileRequest,
) (*fs.DownloadFileResponse, error) {
	file, err := s.svc.DownloadFile(ctx, req.Filename)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.NotFound, "file not found")
	}
	protoFile := model.ToProtoFile(file)
	return &fs.DownloadFileResponse{File: protoFile}, nil
}

func (s *Server) DeleteFile(ctx context.Context, req *fs.DeleteFileRequest,
) (*fs.DeleteFileResponse, error) {
	if err := s.svc.DeleteFile(ctx, req.Filename); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &fs.DeleteFileResponse{}, nil
}
