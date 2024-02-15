package server

import (
	"context"
	"log"

	"github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileSyncService interface {
	UploadFile(ctx context.Context, filename *proto.File) error
	DownloadFile(ctx context.Context, filename string) (*proto.File, error)
	DeleteFile(ctx context.Context, filename string) error
}

type FileSyncServer struct {
	proto.UnimplementedFileSyncServiceServer
	svc FileSyncService
}

func Register(gRPCServer *grpc.Server, service FileSyncService) {
	proto.RegisterFileSyncServiceServer(gRPCServer, &FileSyncServer{
		svc: service,
	})
}

func (s *FileSyncServer) UploadFile(ctx context.Context, req *proto.UploadFileRequest,
) (*proto.UploadFileResponse, error) {
	if err := s.svc.UploadFile(ctx, req.File); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &proto.UploadFileResponse{}, nil
}

func (s *FileSyncServer) DownloadFile(ctx context.Context, req *proto.DownloadFileRequest,
) (*proto.DownloadFileResponse, error) {
	file, err := s.svc.DownloadFile(ctx, req.Filename)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.NotFound, "file not found")
	}
	return &proto.DownloadFileResponse{File: file}, nil
}

func (s *FileSyncServer) DeleteFile(ctx context.Context, req *proto.DeleteFileRequest,
) (*proto.DeleteFileResponse, error) {
	if err := s.svc.DeleteFile(ctx, req.Filename); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &proto.DeleteFileResponse{}, nil
}
