package events

import (
	"encoding/json"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/models"
	"github.com/google/uuid"
)

type EventType string

const (
	EventTypeBookCreated EventType = "BOOK_CREATED"
	EventTypeBookUpdated EventType = "BOOK_UPDATED"
	EventTypeBookDeleted EventType = "BOOK_DELETED"
)

type Event struct {
	ID        string          `json:"id"`
	Type      EventType       `json:"type"`
	Data      json.RawMessage `json:"data"`
	Timestamp time.Time       `json:"timestamp"`
}

type BookEvent struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func NewBookEvent(eventType EventType, book *models.Book) (*Event, error) {
	bookEvent := BookEvent{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	data, err := json.Marshal(bookEvent)
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:        uuid.New().String(),
		Type:      eventType,
		Data:      data,
		Timestamp: time.Now(),
	}, nil
}

func NewBookDeletedEvent(bookID uint) *Event {
	data, _ := json.Marshal(map[string]interface{}{
		"id": bookID,
	})

	return &Event{
		ID:        uuid.New().String(),
		Type:      EventTypeBookDeleted,
		Data:      data,
		Timestamp: time.Now(),
	}
}
