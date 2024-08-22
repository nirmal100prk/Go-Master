package errhandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomError struct {
	Message    string
	Field      string
	StatusCode codes.Code // gRPC status code
}

func (e *CustomError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	return e.Message
}

// NewCustomError creates a new custom error with a specified message and gRPC status code.
func NewCustomError(message string, field string, statusCode codes.Code) *CustomError {
	return &CustomError{
		Message:    message,
		Field:      field,
		StatusCode: statusCode,
	}
}

// HandleServiceerror handles  errors to gRPC status codes with custom messages.
func HandleServiceerror(err error) error {
	if err == nil {
		return nil
	}

	// Handle custom errors
	if customErr, ok := err.(*CustomError); ok {
		return status.Errorf(customErr.StatusCode, customErr.Error())
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// Handle specific PostgreSQL errors with custom messages
		switch pgErr.Code {
		case "23505": // Unique violation
			return status.Errorf(codes.AlreadyExists, "A record with the same %s already exists.", extractUniqueViolationDetail(pgErr))
		case "23503": // Foreign key violation
			return status.Errorf(codes.FailedPrecondition, "Invalid reference: The %s you provided does not exist.", extractForeignKeyViolationDetail(pgErr))
		case "23514": // Check violation
			return status.Errorf(codes.InvalidArgument, "The provided %s does not meet the required conditions.", extractCheckViolationDetail(pgErr))
		case "22P02": // Invalid text representation
			return status.Errorf(codes.InvalidArgument, "The input for %s is in an invalid format.", extractInvalidTextRepresentationDetail(pgErr))
		default:
			return status.Errorf(codes.Internal, "An internal database error occurred.")
		}
	}

	// Handle other errors, such as context cancellations or deadlines
	if errors.Is(err, context.DeadlineExceeded) {
		return status.Errorf(codes.DeadlineExceeded, "The operation timed out.")
	}
	if errors.Is(err, context.Canceled) {
		return status.Errorf(codes.Canceled, "The operation was canceled.")
	}

	// Default to internal error for unknown cases
	return status.Errorf(codes.Internal, "An unexpected error occurred.")
}

// Helper functions to extract details from PgError for different types of violations
func extractUniqueViolationDetail(pgErr *pgconn.PgError) string {
	// Example: "Key (username)=(johndoe) already exists."
	return pgErr.Detail // Or parse and return a specific part if needed
}

func extractForeignKeyViolationDetail(pgErr *pgconn.PgError) string {
	// Example: "Key (department_id)=(123) is not present in table \"departments\"."
	return pgErr.Detail // Or parse and return a specific part if needed
}

func extractCheckViolationDetail(pgErr *pgconn.PgError) string {
	// Example: "New row violates check constraint \"salary_check\"."
	return pgErr.ConstraintName // Or custom message based on the constraint
}

func extractInvalidTextRepresentationDetail(pgErr *pgconn.PgError) string {
	// Example: "invalid input syntax for type integer: \"abc\""
	return pgErr.Message // Or custom message based on the error
}
