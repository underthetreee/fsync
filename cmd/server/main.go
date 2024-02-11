package main

import (
	"log"
	"net"

	"github.com/underthetreee/fsync/internal/server"
	"github.com/underthetreee/fsync/internal/service"
	manager "github.com/underthetreee/fsync/pkg/file_manager"
	"google.golang.org/grpc"
)

func main() {
	m, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewFileSyncService(m)
	gRPCServer := grpc.NewServer()
	server.Register(gRPCServer, svc)
	l, _ := net.Listen("tcp", ":50051")

	log.Println("gRPC server is listening on", l.Addr())
	if err := gRPCServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
