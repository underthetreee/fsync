package main

import (
	"context"
	"log"

	"github.com/underthetreee/fsync/pkg/kafka"
)

const kafkaTopic = "sync"

func main() {
	ctx := context.Background()
	broker := kafka.NewKafkaConsumer(kafkaTopic)
	defer broker.Close()
	for {
		event, err := broker.ConsumeFileEvent(ctx)
		if err != nil {
			log.Println("read message:", err)
		}
		log.Printf("filename: %v, action: %v\n", event.Filename, event.Action)
	}
}
