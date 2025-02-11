package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(cfg *config.Config) (*Consumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.Topic,
		GroupID: "books-api-consumer",
	})

	return &Consumer{
		reader: reader,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				message, err := c.reader.ReadMessage(ctx)
				if err != nil {
					log.Printf("Error reading message: %v\n", err)
					continue
				}

				var event Event
				if err := json.Unmarshal(message.Value, &event); err != nil {
					log.Printf("Error unmarshaling event: %v\n", err)
					continue
				}

				log.Printf("Received event: %s - %s\n", event.Type, string(event.Data))
			}
		}
	}()

	return nil
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
