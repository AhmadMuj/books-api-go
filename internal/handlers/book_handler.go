package handlers

import (
	"net/http"
	"strconv"

	"github.com/AhmadMuj/books-api-go/internal/dto"
	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"github.com/AhmadMuj/books-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Book details"
// @Success 201 {object} dto.BookResponse
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	book := &models.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := h.bookService.CreateBook(c.Request.Context(), book); err != nil {
		statusCode := http.StatusInternalServerError
		if _, ok := err.(*errors.AppError); ok {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, err)
		return
	}

	c.JSON(http.StatusCreated, dto.ToBookResponse(book))
}

// @Summary Get a book by ID
// @Description Get a book's details by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid book ID"))
		return
	}

	book, err := h.bookService.GetBook(c.Request.Context(), uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		if appErr, ok := err.(*errors.AppError); ok {
			if appErr.Type == errors.NotFound {
				statusCode = http.StatusNotFound
			}
		}
		c.JSON(statusCode, err)
		return
	}

	c.JSON(http.StatusOK, dto.ToBookResponse(book))
}

// @Summary List all books
// @Description Get a paginated list of books
// @Tags books
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} dto.ListBooksResponse
// @Failure 500 {object} errors.AppError
// @Router /books [get]
func (h *BookHandler) ListBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	books, total, err := h.bookService.ListBooks(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	totalPages := (int(total) + pageSize - 1) / pageSize

	response := dto.ListBooksResponse{
		Books:      dto.ToBookResponseList(books),
		Page:       page,
		PageSize:   pageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Update a book
// @Description Update a book's details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.UpdateBookRequest true "Book details"
// @Success 200 {object} dto.BookResponse
// @Failure 400 {object} errors.AppError
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid book ID"))
		return
	}

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	book := &models.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := h.bookService.UpdateBook(c.Request.Context(), uint(id), book); err != nil {
		statusCode := http.StatusInternalServerError
		if appErr, ok := err.(*errors.AppError); ok {
			switch appErr.Type {
			case errors.NotFound:
				statusCode = http.StatusNotFound
			case errors.ValidationErr:
				statusCode = http.StatusBadRequest
			}
		}
		c.JSON(statusCode, err)
		return
	}

	c.JSON(http.StatusOK, dto.ToBookResponse(book))
}

// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid book ID"))
		return
	}

	if err := h.bookService.DeleteBook(c.Request.Context(), uint(id)); err != nil {
		statusCode := http.StatusInternalServerError
		if appErr, ok := err.(*errors.AppError); ok {
			if appErr.Type == errors.NotFound {
				statusCode = http.StatusNotFound
			}
		}
		c.JSON(statusCode, err)
		return
	}

	c.Status(http.StatusNoContent)
}
