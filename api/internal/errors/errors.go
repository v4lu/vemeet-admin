package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorType string

const (
	ErrorTypeValidation     ErrorType = "VALIDATION_ERROR"
	ErrorTypeAuthentication ErrorType = "AUTHENTICATION_ERROR"
	ErrorTypeNotFound       ErrorType = "NOT_FOUND"
	ErrorTypeInternal       ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type       ErrorType
	Message    string
	HTTPStatus int
	Raw        error
}

func (e *AppError) Error() string {
	if e.Raw != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Type, e.Message, e.Raw)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeValidation,
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeAuthentication,
		Message:    message,
		HTTPStatus: http.StatusUnauthorized,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeNotFound,
		Message:    message,
		HTTPStatus: http.StatusNotFound,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeInternal,
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
	}
}

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.HTTPStatus, gin.H{
			"error": gin.H{
				"type":    appErr.Type,
				"message": appErr.Message,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": gin.H{
			"type":    ErrorTypeInternal,
			"message": "An unexpected error occurred",
		},
	})
}
