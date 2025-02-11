package service

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/cache"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"github.com/AhmadMuj/books-api-go/internal/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *models.Book) error
	GetBook(ctx context.Context, id uint) (*models.Book, error)
	ListBooks(ctx context.Context, page, pageSize int) ([]models.Book, int64, error)
	UpdateBook(ctx context.Context, id uint, book *models.Book) error
	DeleteBook(ctx context.Context, id uint) error
}

type bookService struct {
	repo  repository.BookRepository
	cache cache.Cache
}

func NewBookService(repo repository.BookRepository, cache cache.Cache) BookService {
	return &bookService{
		repo:  repo,
		cache: cache,
	}
}
