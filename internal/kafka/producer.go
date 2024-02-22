package kafka

import (
	"context"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/protobuf/proto"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(topic string) *KafkaProducer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		Async:                  true,
		AllowAutoTopicCreation: true,
	}
	return &KafkaProducer{
		writer: w,
	}
}

func (p *KafkaProducer) ProduceFileEvent(ctx context.Context, event *fs.FileEvent) error {
	eventBytes, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	if err = p.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(event.GetFilename()),
			Value: eventBytes,
		},
	); err != nil {
		return err
	}
	timestamp := time.Unix(event.Timestamp.Seconds, int64(event.Timestamp.Nanos)).Format(time.RFC822)
	slog.Info("produce event",
		"file", event.Filename,
		"action", event.Action,
		"timestamp", timestamp,
	)
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
