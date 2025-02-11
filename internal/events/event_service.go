package events

import (
	"context"
	"fmt"

	"github.com/AhmadMuj/books-api-go/internal/models"
)

type EventService interface {
	PublishBookCreated(ctx context.Context, book *models.Book) error
	PublishBookUpdated(ctx context.Context, book *models.Book) error
	PublishBookDeleted(ctx context.Context, bookID uint) error
}

type eventService struct {
	producer Producer
}

func NewEventService(producer Producer) EventService {
	return &eventService{
		producer: producer,
	}
}

func (s *eventService) PublishBookCreated(ctx context.Context, book *models.Book) error {
	event, err := NewBookEvent(EventTypeBookCreated, book)
	if err != nil {
		return fmt.Errorf("failed to create book created event: %w", err)
	}

	return s.producer.PublishEvent(ctx, event)
}

func (s *eventService) PublishBookUpdated(ctx context.Context, book *models.Book) error {
	event, err := NewBookEvent(EventTypeBookUpdated, book)
	if err != nil {
		return fmt.Errorf("failed to create book updated event: %w", err)
	}

	return s.producer.PublishEvent(ctx, event)
}

func (s *eventService) PublishBookDeleted(ctx context.Context, bookID uint) error {
	event := NewBookDeletedEvent(bookID)
	return s.producer.PublishEvent(ctx, event)
}
