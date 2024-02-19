package main

import (
	"context"
	"log"

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

	c, err := grpc.NewClient(listenAddr, broker)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.InitEventLoop(ctx); err != nil {
		log.Fatal(err)
	}
}
