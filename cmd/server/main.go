package main

import (
	"log"

	manager "github.com/underthetreee/fsync/internal/file_manager"
	"github.com/underthetreee/fsync/internal/kafka"
	"github.com/underthetreee/fsync/internal/service"
	"github.com/underthetreee/fsync/internal/transport/grpc"
)

const (
	kafkaTopic = "sync"
	listenAddr = ":50051"
)

func main() {
	mng, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewFileSyncService(mng)

	broker := kafka.NewKafkaProducer(kafkaTopic)
	defer broker.Close()

	srv := grpc.NewServer(svc, broker)

	log.Println("grpc server is listening on", listenAddr)
	if err := srv.Run(listenAddr); err != nil {
		log.Fatal(err)
	}
}
