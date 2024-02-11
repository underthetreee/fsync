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
	UploadFile(context.Context, *proto.File) error
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

func (s *FileSyncServer) UploadFile(ctx context.Context, req *proto.UploadFileRequest) (*proto.UploadFileResponse, error) {
	file := req.GetFile()

	if err := s.svc.UploadFile(ctx, file); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &proto.UploadFileResponse{}, nil
}

func (s *FileSyncServer) DownloadFile(ctx context.Context, req *proto.DownloadFileRequest) (*proto.DownloadFileResponse, error) {
	return &proto.DownloadFileResponse{File: nil}, nil
}

func (s *FileSyncServer) DeleteFile(ctx context.Context, req *proto.DeleteFileRequest) (*proto.DeleteFileResponse, error) {
	return &proto.DeleteFileResponse{}, nil
}
