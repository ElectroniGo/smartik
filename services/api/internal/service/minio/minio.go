package minio

import (
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(endpointUrl, accessId, secretKey string) (*minio.Client, error) {
	client, err := minio.New(endpointUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(accessId, secretKey, ""),
		Secure: false,
	})

	return client, err
}
