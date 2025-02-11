package repository

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, id uint) (*models.Book, error)
	List(ctx context.Context, limit, offset int) ([]models.Book, int64, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id uint) error
}
