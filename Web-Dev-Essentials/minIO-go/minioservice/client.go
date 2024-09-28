package minioservice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioServiceClient struct {
	Client *minio.Client
}

func NewMinioService(accessKey, secretKey, minioURL string) (*MinioServiceClient, error) {
	clientInstance, err := minio.New(minioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &MinioServiceClient{
		Client: clientInstance,
	}, nil
}

func (s *MinioServiceClient) UploadObject(ctx context.Context, bucketname, objName string, content []byte, contentType string) error {

	objStats, err := s.Client.StatObject(ctx, bucketname, objName, minio.GetObjectOptions{})
	if err != nil {
		minioErr, ok := err.(minio.ErrorResponse)
		if ok && minioErr.Code != MinioNoSuchKey {
			return handleMinioError(err)
		}

	}
	if objStats.Key != "" {
		return errors.New("object key already exists")
	}
	objectBytes := bytes.NewReader(content)

	_, err = s.Client.PutObject(ctx, bucketname, objName, objectBytes, objectBytes.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return handleMinioError(err)
	}
	return nil
}

func (s *MinioServiceClient) DownloadObject(ctx context.Context, bucketname, objName string) ([]byte, error) {

	objStats, err := s.Client.StatObject(ctx, bucketname, objName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, handleMinioError(err)
	}

	obj, err := s.Client.GetObject(ctx, bucketname, objName, minio.GetObjectOptions{})
	if err != nil {
		return nil, handleMinioError(err)
	}
	// read bytes to output object
	objectBytes := make([]byte, objStats.Size)
	bytesRead := 0
	for {
		n, err := obj.Read(objectBytes)
		bytesRead += n
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				return nil, err
			}
		}
	}
	if bytesRead != int(objStats.Size) {
		return nil, errors.New("Could not get object")
	}
	return objectBytes, nil
}

func (s *MinioServiceClient) DeleteObject(ctx context.Context, bucketname, objName string) error {

	err := s.Client.RemoveObject(ctx, bucketname, objName, minio.RemoveObjectOptions{})
	if err != nil {
		return handleMinioError(err)
	}
	return nil
}
