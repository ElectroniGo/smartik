package minio

import (
	"context"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/smartik/api/internal/config"
)

func NewMinioClient(endpointUrl, accessId, secretKey string) (*minio.Client, error) {
	client, err := minio.New(endpointUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(accessId, secretKey, ""),
		Secure: false,
	})

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	ok, err := client.BucketExists(context.Background(), cfg.MinioStorageBucket)
	if err != nil {
		return nil, err
	}

	// Create new bucket if it does not exist
	if !ok {
		err := client.MakeBucket(context.Background(), cfg.MinioStorageBucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return client, err
}
