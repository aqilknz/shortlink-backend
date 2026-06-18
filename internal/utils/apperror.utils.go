package utils

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	// 400 Bad Request
	ErrInvalidInput      = NewAppError(http.StatusBadRequest, "Invalid input data format")
	ErrInvalidEmail      = NewAppError(http.StatusBadRequest, "Invalid email format")
	ErrInvalidURL        = NewAppError(http.StatusBadRequest, "The provided URL is invalid")
	ErrInvalidSlugLength = NewAppError(http.StatusBadRequest, "Custom slug must be between 3 and 50 characters")
	ErrInvalidSlugFormat = NewAppError(http.StatusBadRequest, "Custom slug can only contain letters, numbers, and hyphens")
	ErrReservedSlug      = NewAppError(http.StatusBadRequest, "This custom slug is reserved and cannot be used")

	// 401 Unauthorized (Gagal Autentikasi / Token Invalid)
	ErrInvalidCredentials = NewAppError(http.StatusUnauthorized, "Invalid email or password")
	ErrUnauthorized       = NewAppError(http.StatusUnauthorized, "Invalid or expired session, please login again")

	// 403 Forbidden (Tidak Punya Akses)
	ErrForbidden = NewAppError(http.StatusForbidden, "You do not have permission to access or modify this resource")

	// 404 Not Found
	ErrNotFound = NewAppError(http.StatusNotFound, "Resource not found")

	// 409 Conflict
	ErrEmailAlreadyExists = NewAppError(http.StatusConflict, "Email is already registered")
	ErrSlugAlreadyExists  = NewAppError(http.StatusConflict, "Short link slug is already in use")

	// 422 Unprocessable Entity
	ErrValidationFailed = NewAppError(http.StatusUnprocessableEntity, "Data validation failed")

	// 429 Too Many Requests
	ErrTooManyRequests = NewAppError(http.StatusTooManyRequests, "Too many requests, please try again later")

	// 500 Internal Server Error
	ErrInternalServer = NewAppError(http.StatusInternalServerError, "Internal server error")
)
