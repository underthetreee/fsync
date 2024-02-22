package main

import (
	"context"
	"log"

	manager "github.com/underthetreee/fsync/internal/file_manager"
	"github.com/underthetreee/fsync/internal/kafka"
	"github.com/underthetreee/fsync/internal/transport/grpc"
)

const (
	kafkaTopic = "sync"
	listenAddr = ":50051"
)

func main() {
	ctx := context.Background()

	broker := kafka.NewKafkaConsumer(kafkaTopic)
	defer broker.Close()

	mng, err := manager.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	c, err := grpc.NewClient(listenAddr, broker, mng)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.SyncLoop(ctx); err != nil {
		log.Println(err)
	}
}
