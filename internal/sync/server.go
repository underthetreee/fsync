package sync

import (
	"context"

	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/grpc"
)

type FileSyncServer struct {
	fs.UnimplementedFileSyncServiceServer
}

func Register(gRPCServer *grpc.Server) {
	fs.RegisterFileSyncServiceServer(gRPCServer, &FileSyncServer{})
}

func (s *FileSyncServer) UploadFile(ctx context.Context, req *fs.UploadFileRequest) (*fs.UploadFileResponse, error) {
	return &fs.UploadFileResponse{
		Success: true,
		Message: "file uploaded",
	}, nil
}
func (s *FileSyncServer) DownloadFile(ctx context.Context, req *fs.DownloadFileRequest) (*fs.DownloadFileResponse, error) {
	return &fs.DownloadFileResponse{
		File:    nil,
		Success: true,
		Message: "file downloaded",
	}, nil
}
func (s *FileSyncServer) DeleteFile(ctx context.Context, req *fs.DeleteFileRequest) (*fs.DeleteFileResponse, error) {
	return &fs.DeleteFileResponse{
		Success: true,
		Message: "file deleted",
	}, nil
}
