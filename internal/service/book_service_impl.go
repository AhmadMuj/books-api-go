package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
)

func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	if err := validateBook(book); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, book); err != nil {
		return err
	}

	// Invalidate list cache
	if err := s.cache.InvalidateBooksList(ctx); err != nil {
		fmt.Printf("Failed to invalidate books list cache: %v\n", err)
	}

	if err := s.eventService.PublishBookCreated(ctx, book); err != nil {
		log.Printf("Failed to publish book created event: %v\n", err)
	}

	return nil
}

func (s *bookService) GetBook(ctx context.Context, id uint) (*models.Book, error) {
	if id == 0 {
		return nil, errors.NewValidationError("invalid book ID")
	}

	// Try to get from cache first
	if book, err := s.cache.GetBook(ctx, id); err == nil && book != nil {
		return book, nil
	}

	// If not in cache, get from database
	book, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Cache the book for future requests
	if book != nil {
		if err := s.cache.SetBook(ctx, book); err != nil {
			// Log error but don't fail the request
			fmt.Printf("Failed to cache book: %v\n", err)
		}
	}

	return book, nil
}

func (s *bookService) ListBooks(ctx context.Context, page, pageSize int) ([]models.Book, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Try to get from cache first
	if books, total, err := s.cache.GetBooksList(ctx, page, pageSize); err == nil && books != nil {
		return books, total, nil
	}

	// If not in cache, get from database
	books, total, err := s.repo.List(ctx, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}

	// Cache the results
	if err := s.cache.SetBooksList(ctx, books, total, page, pageSize); err != nil {
		fmt.Printf("Failed to cache books list: %v\n", err)
	}

	return books, total, nil
}

func (s *bookService) UpdateBook(ctx context.Context, id uint, book *models.Book) error {
	if err := s.repo.Update(ctx, book); err != nil {
		return err
	}

	// Invalidate both single book and list caches
	if err := s.cache.DeleteBook(ctx, id); err != nil {
		fmt.Printf("Failed to invalidate book cache: %v\n", err)
	}
	if err := s.cache.InvalidateBooksList(ctx); err != nil {
		fmt.Printf("Failed to invalidate books list cache: %v\n", err)
	}
	if err := s.eventService.PublishBookUpdated(ctx, book); err != nil {
		log.Printf("Failed to publish book created event: %v\n", err)
	}

	return nil
}

func (s *bookService) DeleteBook(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	// Invalidate both single book and list caches
	if err := s.cache.DeleteBook(ctx, id); err != nil {
		fmt.Printf("Failed to invalidate book cache: %v\n", err)
	}
	if err := s.cache.InvalidateBooksList(ctx); err != nil {
		fmt.Printf("Failed to invalidate books list cache: %v\n", err)
	}
	if err := s.eventService.PublishBookDeleted(ctx, id); err != nil {
		log.Printf("Failed to publish book created event: %v\n", err)
	}

	return nil
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
