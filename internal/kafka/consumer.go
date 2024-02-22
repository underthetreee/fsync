package kafka

import (
	"context"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
	fs "github.com/underthetreee/fsync/pkg/proto"
	"google.golang.org/protobuf/proto"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(topic string) *KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: "events",
	})
	return &KafkaConsumer{
		reader: r,
	}
}

func (c *KafkaConsumer) ConsumeFileEvent(ctx context.Context) (*fs.FileEvent, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}
	event := &fs.FileEvent{}
	if err := proto.Unmarshal(msg.Value, event); err != nil {
		return nil, err
	}
	timestamp := time.Unix(event.Timestamp.Seconds, int64(event.Timestamp.Nanos)).Format(time.RFC822)
	slog.Info("consume event",
		"file", event.Filename,
		"action", event.Action,
		"timestamp", timestamp,
	)
	return event, nil
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
