package events

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishEvent(ctx context.Context, event *Event) error
	Close() error
}

type KafkaProducer struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaProducer(cfg *config.Config) (Producer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.Topic,
		// Async production for better performance
		Async: true,
	})

	return &KafkaProducer{
		writer: writer,
		topic:  cfg.Kafka.Topic,
	}, nil
}

func (p *KafkaProducer) PublishEvent(ctx context.Context, event *Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	message := kafka.Message{
		Key:   []byte(event.Type),
		Value: data,
	}

	if err := p.writer.WriteMessages(ctx, message); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
