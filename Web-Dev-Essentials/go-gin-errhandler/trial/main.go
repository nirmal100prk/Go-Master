package main

import (
	"fmt"
	"net/http"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CustomError structure to hold error information
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

// Implementing the Error method to satisfy the error interface
func (r *CustomError) Error() string {
	return r.Message
}

// Implementing the Stringer interface to customize the fmt.Println output
func (r *CustomError) String() string {
	return fmt.Sprintf("Code: %d, Message: %s,  Original Error: %v",
		r.Code, r.Message, r.Err)
}

// Helper function to update struct fields using reflection
func updateStructFields(obj any, customErr *CustomError) {
	v := reflect.ValueOf(obj)

	// Ensure we're working with a pointer to a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Get the value that the pointer points to
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Get the value that the pointer points to
	}
	fmt.Println(v.Kind())
	if v.Kind() == reflect.Struct {
		// Update the Code field if it exists
		codeField := v.FieldByName("Code")
		if codeField.IsValid() && codeField.CanSet() {
			if codeField.Kind() == reflect.Int32 || codeField.Kind() == reflect.Int || codeField.Kind() == reflect.Int64 {
				codeField.SetInt(int64(customErr.Code))
			}
		}

		// Update the Message field if it exists
		messageField := v.FieldByName("Message")
		if messageField.IsValid() && messageField.CanSet() {
			if messageField.Kind() == reflect.String {
				messageField.SetString(customErr.Message)
			}
		}
	}
}

// grpcToHTTPStatus converts gRPC status codes to HTTP status codes
func grpcToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

// NewCustomError returns an instance of CustomError and updates responseObj with code and message
func InjectErrorIntoResponse(customErr *CustomError, responseObj any, err error) *CustomError {
	// Check if the error is a gRPC error
	st, ok := status.FromError(err)
	if ok && st != nil && st.Code() != codes.OK {
		// Extract gRPC status code and message
		customErr.Code = grpcToHTTPStatus(st.Code()) // Convert gRPC code to HTTP code
		customErr.Message = st.Message()
	}
	if responseObj == nil {
		responseObj = make(map[string]interface{})
	}

	// Check if responseObj is a map
	if respMap, ok := responseObj.(map[string]interface{}); ok {
		// If it's a map, update or add the code and message fields

		respMap["Code"] = customErr.Code
		respMap["Message"] = customErr.Message
	} else if reflect.TypeOf(responseObj).Kind() == reflect.Struct || reflect.TypeOf(responseObj).Kind() == reflect.Ptr {
		// If responseObj is a struct or a pointer to a struct, update fields dynamically
		updateStructFields(&responseObj, customErr)
	} else {
		// If responseObj is neither a map nor a struct, initialize it as a map and set code and message
		responseObj = map[string]interface{}{
			"code":    customErr.Code,
			"message": customErr.Message,
		}
	}

	return &CustomError{}
}

func main() {
	// Example with a non-gRPC error (passing nil for the error)
	// respMap := map[string]interface{}{
	// 	"Code":    200,
	// 	"Message": "success",
	// 	"Other":   "Some other data",
	// }
	// RespStruct := struct {
	// 	Code    int32
	// 	Message string
	// 	Other   string
	// }{
	// 	Other: "Some other data",
	// }
	// cObjWithMap := NewCustomError(0, "Invalid Request", "Something is wrong with the input", &RespStruct, nil)
	// fmt.Println(cObjWithMap.String())

	// Example with a gRPC error
	// grpcErr := status.Error(codes.InvalidArgument, "Invalid argument provided")
	// cObjWithGRPC := NewCustomError(0, "", "gRPC error", &RespStruct, grpcErr)
	// fmt.Println(cObjWithGRPC.String())
}
