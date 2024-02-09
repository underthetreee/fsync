package main

import (
	"log"
	"net"

	"github.com/underthetreee/fsync/internal/model"
	"github.com/underthetreee/fsync/internal/sync"
	manager "github.com/underthetreee/fsync/pkg/file_manager"
	"google.golang.org/grpc"
)

func main() {
	m, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	file := &model.File{
		Filename: "test_file.txt",
		Content:  []byte("this is the some1\n"),
	}
	if err := m.CreateFile(file); err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	sync.Register(gRPCServer)
	l, _ := net.Listen("tcp", ":50051")

	log.Println("gRPC server is listening on", l.Addr())
	if err := gRPCServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
