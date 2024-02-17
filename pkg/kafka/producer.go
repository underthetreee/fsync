package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/underthetreee/fsync/internal/model"
	"google.golang.org/protobuf/proto"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	w := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
	}
	return &KafkaProducer{
		writer: w,
	}
}

func (p *KafkaProducer) ProduceFileEvent(ctx context.Context, topic string, event *model.FileEvent) error {
	protoEvent := model.ToProtoEvent(event)
	eventBytes, err := proto.Marshal(protoEvent)
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
	return nil
}
