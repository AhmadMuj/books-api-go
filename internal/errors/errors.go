package errors

import "fmt"

type ErrorType string

const (
	NotFound      ErrorType = "NOT_FOUND"
	AlreadyExists ErrorType = "ALREADY_EXISTS"
	ValidationErr ErrorType = "VALIDATION_ERROR"
	DatabaseErr   ErrorType = "DATABASE_ERROR"
	InternalErr   ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Constructor functions for common errors
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
	}
}

func NewAlreadyExistsError(message string) *AppError {
	return &AppError{
		Type:    AlreadyExists,
		Message: message,
	}
}

func NewDatabaseError(err error) *AppError {
	return &AppError{
		Type:    DatabaseErr,
		Message: "database operation failed",
		Err:     err,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ValidationErr,
		Message: message,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Type:    InternalErr,
		Message: "internal server error",
		Err:     err,
	}
}
