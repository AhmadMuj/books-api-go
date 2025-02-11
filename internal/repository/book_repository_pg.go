package repository

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"gorm.io/gorm"
)

type BookRepositoryPG struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryPG{
		db: db,
	}
}

func (r *BookRepositoryPG) Create(ctx context.Context, book *models.Book) error {
	// Check if book with same title and author exists
	var exists bool
	err := r.db.WithContext(ctx).
		Model(&models.Book{}).
		Select("count(*) > 0").
		Where("title = ? AND author = ?", book.Title, book.Author).
		Find(&exists).
		Error
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	if exists {
		return errors.NewAlreadyExistsError("book with same title and author already exists")
	}

	result := r.db.WithContext(ctx).Create(book)
	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}
	return nil
}

func (r *BookRepositoryPG) GetByID(ctx context.Context, id uint) (*models.Book, error) {
	var book models.Book
	result := r.db.WithContext(ctx).First(&book, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("book not found")
		}
		return nil, errors.NewDatabaseError(result.Error)
	}
	return &book, nil
}

func (r *BookRepositoryPG) List(ctx context.Context, limit, offset int) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	// Get total count
	if err := r.db.WithContext(ctx).Model(&models.Book{}).Count(&total).Error; err != nil {
		return nil, 0, errors.NewDatabaseError(err)
	}

	result := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&books)

	if result.Error != nil {
		return nil, 0, errors.NewDatabaseError(result.Error)
	}
	return books, total, nil
}

func (r *BookRepositoryPG) Update(ctx context.Context, book *models.Book) error {
	result := r.db.WithContext(ctx).Save(book)
	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("book not found")
	}
	return nil
}

func (r *BookRepositoryPG) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Book{}, id)
	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("book not found")
	}
	return nil
}
