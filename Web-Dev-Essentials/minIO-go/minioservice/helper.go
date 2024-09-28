package minioservice

import (
	"log/slog"

	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	MinioNoSuchBucket string = "NoSuchBucket"
	MinioNoSuchKey    string = "NoSuchKey"
	MinioAccessDenied string = "AccessDenied"
)

func handleMinioError(err error) error {
	if err != nil {
		slog.Error("MinIO Error: ", slog.String("error", err.Error()))
	} else {
		return nil
	}
	switch t := err.(type) {
	case minio.ErrorResponse:
		if t.Code == MinioNoSuchBucket {
			return status.Errorf(codes.NotFound, "Bucket does not exist")
		} else if t.Code == MinioNoSuchKey {
			return status.Errorf(codes.NotFound, "Object with the specified key is not found")
		} else if t.Code == MinioAccessDenied {
			return status.Errorf(codes.Unauthenticated, "Access Denied")
		} else {
			return status.Errorf(codes.Unknown, "Internal Server Error")
		}
	default:
		return status.Errorf(codes.Unknown, "Internal Server Error")
	}
}
