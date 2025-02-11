package cache

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/models"
)

type Cache interface {
	// Single book operations
	GetBook(ctx context.Context, id uint) (*models.Book, error)
	SetBook(ctx context.Context, book *models.Book) error
	DeleteBook(ctx context.Context, id uint) error

	// Book list operations
	GetBooksList(ctx context.Context, page, pageSize int) ([]models.Book, int64, error)
	SetBooksList(ctx context.Context, books []models.Book, total int64, page, pageSize int) error
	InvalidateBooksList(ctx context.Context) error

	// Optional: General cache operations
	Clear(ctx context.Context) error
	Close() error
}
