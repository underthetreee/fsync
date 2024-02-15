package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/protobuf/proto"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Balancer: &kafka.LeastBytes{},
	}
	return &KafkaProducer{
		writer: w,
	}
}

func (p *KafkaProducer) ProduceFileEvent(ctx context.Context, topic string, event *fs.FileEvent) error {
	eventBytes, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	if err = p.writer.WriteMessages(ctx,
		kafka.Message{
			Topic: topic,
			Key:   []byte(event.Filename),
			Value: eventBytes,
		},
	); err != nil {
		return err

	}
	log.Println("event produced", eventBytes)
	return nil
}
