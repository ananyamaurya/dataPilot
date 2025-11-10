package storage

import (
    "context"
    "fmt"
    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "io"
)

type MinioClient struct {
    client *minio.Client
    bucket string
}

func New(endpoint, accessKey, secretKey, bucket string) (*MinioClient, error) {
    mc, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false, // use true with TLS
    })
    if err != nil {
        return nil, err
    }
    // ensure bucket exists
    ctx := context.Background()
    exists, err := mc.BucketExists(ctx, bucket)
    if err != nil {
        return nil, err
    }
    if !exists {
        err = mc.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
        if err != nil {
            return nil, err
        }
    }
    return &MinioClient{client: mc, bucket: bucket}, nil
}

func (m *MinioClient) Upload(ctx context.Context, objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
    _, err := m.client.PutObject(ctx, m.bucket, objectKey, reader, size, minio.PutObjectOptions{ContentType: contentType})
    if err != nil {
        return "", err
    }
    // return object URL or object key
    return fmt.Sprintf("%s/%s/%s", m.client.EndpointURL(), m.bucket, objectKey), nil
}
