package main

import (
	"log"
	"net"

	manager "github.com/underthetreee/fsync/internal/file_manager"
	"github.com/underthetreee/fsync/internal/server"
	"github.com/underthetreee/fsync/internal/service"
	"github.com/underthetreee/fsync/pkg/kafka"
	"google.golang.org/grpc"
)

func main() {
	mng, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	broker := kafka.NewKafkaProducer()
	svc := service.NewFileSyncService(mng, broker)
	srv := grpc.NewServer()
	server.Register(srv, svc)
	l, _ := net.Listen("tcp", ":50051")

	log.Println("gRPC server is listening on", l.Addr())
	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
