package grpc

import (
	"context"
	"log"
	"net"

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

type EventProducer interface {
	ProduceFileEvent(ctx context.Context, event *fs.FileEvent) error
}

type Server struct {
	fs.UnimplementedFileSyncServiceServer
	srv  *grpc.Server
	svc  FileSyncService
	prod EventProducer
}

func NewServer(service FileSyncService, producer EventProducer) *Server {
	grpcServer := grpc.NewServer()
	s := &Server{
		srv:  grpcServer,
		svc:  service,
		prod: producer,
	}
	fs.RegisterFileSyncServiceServer(grpcServer, s)
	return s
}

func (s *Server) UploadFile(ctx context.Context, req *fs.UploadFileRequest,
) (*fs.UploadFileResponse, error) {
	file := model.ToModel(req.GetFile())

	if err := s.svc.UploadFile(ctx, file); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	event := model.NewFileEvent(file.Filename, model.UPLOAD)

	if err := s.prod.ProduceFileEvent(ctx, event); err != nil {
		return nil, err
	}
	return &fs.UploadFileResponse{}, nil
}

func (s *Server) DownloadFile(ctx context.Context, req *fs.DownloadFileRequest,
) (*fs.DownloadFileResponse, error) {
	filename := req.GetFilename()
	file, err := s.svc.DownloadFile(ctx, filename)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.NotFound, "file not found")
	}
	protoFile := model.ToProto(file)
	return &fs.DownloadFileResponse{File: protoFile}, nil
}

func (s *Server) DeleteFile(ctx context.Context, req *fs.DeleteFileRequest,
) (*fs.DeleteFileResponse, error) {
	filename := req.GetFilename()
	if err := s.svc.DeleteFile(ctx, filename); err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	event := model.NewFileEvent(filename, model.DELETE)
	if err := s.prod.ProduceFileEvent(ctx, event); err != nil {
		return nil, err
	}
	return &fs.DeleteFileResponse{}, nil
}

func (s *Server) Run(listenAddr string) error {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	return s.srv.Serve(l)
}
