package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
)

func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	if err := validateBook(book); err != nil {
		return err
	}
	return s.repo.Create(ctx, book)
}

func (s *bookService) GetBook(ctx context.Context, id uint) (*models.Book, error) {
	if id == 0 {
		return nil, errors.NewValidationError("invalid book ID")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *bookService) ListBooks(ctx context.Context, page, pageSize int) ([]models.Book, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return s.repo.List(ctx, pageSize, offset)
}

func (s *bookService) UpdateBook(ctx context.Context, id uint, book *models.Book) error {
	if id == 0 {
		return errors.NewValidationError("invalid book ID")
	}

	if err := validateBook(book); err != nil {
		return err
	}

	existingBook, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.NewNotFoundError("book not found")
	}

	book.ID = id
	return s.repo.Update(ctx, book)
}

func (s *bookService) DeleteBook(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.NewValidationError("invalid book ID")
	}
	return s.repo.Delete(ctx, id)
}

func validateBook(book *models.Book) error {
	if book == nil {
		return errors.NewValidationError("book cannot be nil")
	}
	if book.Title == "" {
		return errors.NewValidationError("book title is required")
	}
	if book.Author == "" {
		return errors.NewValidationError("book author is required")
	}

	currentYear := time.Now().Year()
	if book.Year < 1500 || book.Year > currentYear {
		return errors.NewValidationError(fmt.Sprintf(
			"book year must be between 1500 and %d",
			currentYear,
		))
	}
	return nil
}
