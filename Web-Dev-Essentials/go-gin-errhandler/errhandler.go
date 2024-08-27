package errormiddleware

import (
	"encoding/json"
	"go-datastructures/Web-Dev-Essentials/go-gin-errhandler/constants"
	"net/http"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorList struct {
	Errors []ErrorResponse `json:"errors"`
}

type CustomError struct {
	Code             int    `json:"code"`
	Message          string `json:"message"`
	CustomErrMessage string `json:"customerrormessage"`
	Err              error
}

func (r *CustomError) Error() string {
	return r.CustomErrMessage
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// centralized function to map errors to http errors
func HandleError(c *gin.Context, errType gin.ErrorType) {
	var errList ErrorList
	detectedErrors := c.Errors.ByType(errType)
	if len(detectedErrors) > 0 {
		err := detectedErrors[0].Err
		slog.Error("error", err)

		isDebugEnabled := c.GetHeader("-x-debug-mode-")
		var errorRes *ErrorResponse

		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			errorRes = &ErrorResponse{Code: http.StatusBadRequest, Message: constants.InvalidJson}
		case *pgconn.PgError:
			errorRes = handlePgxError(e)
		case *CustomError:
			errorRes = &ErrorResponse{Code: e.Code, Message: e.Message}
		default:
			errorRes = SetCustomErrorDetails(err, err.Error())
		}

		if isDebugEnabled == "true" && err.Error() != "" {
			errorRes.Details = err.Error()
		}
		errList.Errors = append(errList.Errors, *errorRes)
		c.AbortWithStatusJSON(errorRes.Code, errList)
		return
	}
}

func handlePgxError(err *pgconn.PgError) *ErrorResponse {
	switch err.Code {
	case constants.UNIQUE_VIOLATION:
		return &ErrorResponse{Code: http.StatusConflict, Message: constants.UniqueViolation}
	case constants.FOREIGNKEY_VIOLATION:
		return &ErrorResponse{Code: http.StatusBadRequest, Message: constants.ForeignKeyViolation}
	case constants.CHECK_VIOLATION:
		return &ErrorResponse{Code: http.StatusBadRequest, Message: constants.CheckViolation}
	default:
		return &ErrorResponse{Code: http.StatusInternalServerError, Message: constants.DBError, Details: err.Message}
	}
}

func SetCustomErrorDetails(err error, errDetails string) *ErrorResponse {
	var errorRes *ErrorResponse
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.DeadlineExceeded:
			errorRes = &ErrorResponse{Code: http.StatusServiceUnavailable, Message: constants.ServiceUnavaibleError}
		case codes.Unavailable:
			errorRes = &ErrorResponse{Code: http.StatusServiceUnavailable, Message: constants.ServiceUnavaibleError}
		case codes.Unauthenticated:
			errorRes = &ErrorResponse{Code: http.StatusUnauthorized, Message: e.Message()}
		case codes.AlreadyExists:
			errorRes = &ErrorResponse{Code: http.StatusConflict, Message: e.Message()}
		case codes.ResourceExhausted:
			errorRes = &ErrorResponse{Code: http.StatusServiceUnavailable, Message: constants.ServiceUnavaibleError}
		case codes.InvalidArgument:
			errorRes = &ErrorResponse{Code: http.StatusBadRequest, Message: constants.InvalidArgument}
		case codes.Internal:
			errorRes = &ErrorResponse{Code: http.StatusInternalServerError, Message: constants.InternalServerError}
		case codes.FailedPrecondition:
			errorRes = &ErrorResponse{Code: http.StatusPreconditionFailed, Message: e.Message()}
		case codes.Unimplemented:
			errorRes = &ErrorResponse{Code: http.StatusNotImplemented, Message: constants.Unimplemented}
		default:
			errorRes = &ErrorResponse{Code: http.StatusInternalServerError, Message: constants.InternalServerError}
		}
		return errorRes
	}
	return &ErrorResponse{Code: http.StatusInternalServerError, Message: constants.InternalServerError, Details: errDetails}
}

// HandleGrpcError is the centralized function to map errors to gRPC status codes.
func HandleGrpcError(err error) error {
	switch e := err.(type) {
	case *pgconn.PgError:
		return handlePgxErrorGrpc(e)
	case *CustomError:
		return status.Error(codes.Code(e.Code), e.Message)
	default:
		return status.Error(codes.Internal, constants.InternalServerError)
	}
}

// handlePgxErrorGrpc maps pgx errors to gRPC status codes.
func handlePgxErrorGrpc(err *pgconn.PgError) error {
	switch err.Code {
	case constants.UNIQUE_VIOLATION:
		return status.Error(codes.AlreadyExists, constants.UniqueViolation)
	case constants.FOREIGNKEY_VIOLATION:
		return status.Error(codes.InvalidArgument, constants.ForeignKeyViolation)
	case constants.CHECK_VIOLATION:
		return status.Error(codes.InvalidArgument, constants.CheckViolation)
	default:
		return status.Error(codes.Internal, err.Message)
	}
}
