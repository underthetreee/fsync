package kafka

import (
	"context"
	"log/slog"

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
	slog.Info("consume event",
		"file", event.Filename,
		"action", event.Action,
	)
	return event, nil
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}
