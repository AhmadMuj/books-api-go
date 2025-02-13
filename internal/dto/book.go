package dto

import (
	"fmt"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
)

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,min=1500"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,min=1500"`
}

type BookResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListBooksResponse struct {
	Books      []BookResponse `json:"books"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalItems int64          `json:"total_items"`
	TotalPages int            `json:"total_pages"`
}

// Conversion helpers
func ToBookResponse(book *models.Book) *BookResponse {
	return &BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Year:      book.Year,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func ToBookResponseList(books []models.Book) []BookResponse {
	responses := make([]BookResponse, len(books))
	for i, book := range books {
		responses[i] = *ToBookResponse(&book)
	}
	return responses
}

// Validation helpers
func (r *CreateBookRequest) Validate() error {
	currentYear := time.Now().Year()
	if r.Year < 1500 || r.Year > currentYear {
		return errors.NewValidationError(fmt.Sprintf(
			"book year must be between 1500 and %d",
			currentYear,
		))
	}
	return nil
}

func (r *UpdateBookRequest) Validate() error {
	currentYear := time.Now().Year()
	if r.Year < 1500 || r.Year > currentYear {
		return errors.NewValidationError(fmt.Sprintf(
			"book year must be between 1500 and %d",
			currentYear,
		))
	}
	return nil
}
