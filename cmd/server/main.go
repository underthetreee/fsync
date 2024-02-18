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

const kafkaTopic = "sync"

func main() {
	mng, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	broker := kafka.NewKafkaProducer(kafkaTopic)
	defer broker.Close()

	svc := service.NewFileSyncService(mng)
	srv := grpc.NewServer()
	server.Register(srv, svc, broker)
	l, _ := net.Listen("tcp", ":50051")

	log.Println("gRPC server is listening on", l.Addr())
	if err := srv.Serve(l); err != nil {
		log.Fatal(err)
	}
}
