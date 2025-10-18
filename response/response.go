package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Response messages
const (
	MessageCreated = "created"
	MessageSuccess = "success"
	MessageUpdated = "updated"
	MessageDeleted = "deleted"
)

// BaseResponse represents a standard API response structure.
type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse represents an error response structure.
type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

// IdResponse represents a response containing an ID.
type IdResponse struct {
	Id string `json:"id"`
}

// PaginationResponse represents a paginated response structure.
type PaginationResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Total   int64  `json:"total"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

// Error mapping configuration for common database and application errors
type errorMapping struct {
	code codes.Code
	msg  string
}

var errorMappings = map[string]errorMapping{
	// Database errors
	"no rows in result set":                          {codes.NotFound, "record not found"},
	"duplicate key value violates unique constraint": {codes.AlreadyExists, "record already exists"},
	"violates foreign key constraint":                {codes.InvalidArgument, "invalid reference to related record"},
	"no rows affected":                               {codes.NotFound, "record not found"},
	"connection refused":                             {codes.Unavailable, "database connection failed"},
	"context deadline exceeded":                      {codes.DeadlineExceeded, "operation timeout"},

	// Validation errors
	"invalid input":  {codes.InvalidArgument, "invalid input data"},
	"required field": {codes.InvalidArgument, "required field is missing"},
	"invalid format": {codes.InvalidArgument, "invalid data format"},

	// Authentication/Authorization errors
	"unauthorized":      {codes.Unauthenticated, "authentication required"},
	"token expired":     {codes.Unauthenticated, "authentication token expired"},
	"invalid token":     {codes.Unauthenticated, "invalid authentication token"},
	"permission denied": {codes.PermissionDenied, "insufficient permissions"},
	"access denied":     {codes.PermissionDenied, "access denied"},
}

// grpcToHTTPStatusCode maps gRPC status codes to HTTP status codes.
var grpcToHTTPStatusCode = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusRequestTimeout,
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Unauthenticated:    http.StatusUnauthorized,
}

// ServiceError converts an error into a gRPC status error with appropriate code.
// This should be used in the service layer to wrap errors before returning them.
//
// Parameters:
//   - err: the original error
//   - code: the gRPC status code (use codes.OK to auto-detect from error message)
//
// Returns a gRPC status error with appropriate code and message.
func ServiceError(err error, code codes.Code) error {
	if err == nil {
		return nil
	}

	// If error is already a status error, return it as is
	if _, ok := status.FromError(err); ok {
		return err
	}

	errMsg := err.Error()
	lowerErrMsg := strings.ToLower(errMsg)

	// Try to match error message with known patterns
	for pattern, mapping := range errorMappings {
		if strings.Contains(lowerErrMsg, pattern) {
			return status.Error(mapping.code, mapping.msg)
		}
	}

	// Use provided code if not OK
	if code != codes.OK {
		return status.Error(code, errMsg)
	}

	// Default to Internal error
	return status.Error(codes.Internal, errMsg)
}

// NewServiceError creates a new gRPC status error with the specified code and message.
// Use this for creating custom service-layer errors.
func NewServiceError(code codes.Code, message string) error {
	return status.Error(code, message)
}

// Success sends a successful response with optional data.
func Success(c *gin.Context, message string, data any) {
	response := BaseResponse{
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

// Created sends a 201 Created response with optional data.
func Created(c *gin.Context, data any) {
	response := BaseResponse{
		Message: MessageCreated,
		Data:    data,
	}
	c.JSON(http.StatusCreated, response)
}

// Paginated sends a paginated response.
func Paginated(c *gin.Context, data any, total int64, page, limit int) {
	response := PaginationResponse{
		Message: MessageSuccess,
		Data:    data,
		Total:   total,
		Page:    page,
		Limit:   limit,
	}
	c.JSON(http.StatusOK, response)
}

// Error sends an error response with the specified HTTP status code.
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Message: message,
	})
}

// ErrorWithCode sends an error response with both HTTP status and a custom error code.
func ErrorWithCode(c *gin.Context, statusCode int, message, code string) {
	c.JSON(statusCode, ErrorResponse{
		Message: message,
		Code:    code,
	})
}

// FromError converts a gRPC status error into an HTTP response.
// This should be used in the handler layer to convert service errors.
func FromError(c *gin.Context, err error) {
	if err == nil {
		Success(c, MessageSuccess, nil)
		return
	}

	// Extract gRPC status from error
	st, ok := status.FromError(err)
	if !ok {
		// If not a gRPC status error, treat as internal error
		Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Get HTTP status code from gRPC code
	httpStatus, exists := grpcToHTTPStatusCode[st.Code()]
	if !exists {
		httpStatus = http.StatusInternalServerError
	}

	// Send error response
	ErrorWithCode(c, httpStatus, st.Message(), st.Code().String())
}

// NotFound is a convenience function for 404 errors.
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "resource not found"
	}
	Error(c, http.StatusNotFound, message)
}

// BadRequest is a convenience function for 400 errors.
func BadRequest(c *gin.Context, message string) {
	if message == "" {
		message = "invalid request"
	}
	Error(c, http.StatusBadRequest, message)
}

// Unauthorized is a convenience function for 401 errors.
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "authentication required"
	}
	Error(c, http.StatusUnauthorized, message)
}

// Forbidden is a convenience function for 403 errors.
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "access forbidden"
	}
	Error(c, http.StatusForbidden, message)
}

// InternalError is a convenience function for 500 errors.
func InternalError(c *gin.Context, message string) {
	if message == "" {
		message = "internal server error"
	}
	Error(c, http.StatusInternalServerError, message)
}

// Conflict is a convenience function for 409 errors.
func Conflict(c *gin.Context, message string) {
	if message == "" {
		message = "resource conflict"
	}
	Error(c, http.StatusConflict, message)
}
