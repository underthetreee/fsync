package main

import (
	"log"
	"net"

	manager "github.com/underthetreee/fsync/internal/file_manager"
	"github.com/underthetreee/fsync/internal/service"
	"github.com/underthetreee/fsync/internal/sync"
	"google.golang.org/grpc"
)

func main() {
	mng, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewFileSyncService(mng)
	srv := grpc.NewServer()
	sync.Register(srv, svc)
	l, _ := net.Listen("tcp", ":50051")

	log.Println("gRPC server is listening on", l.Addr())
	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
