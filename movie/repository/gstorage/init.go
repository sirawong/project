package gstorage

import (
	"context"
	"mime/multipart"

	"movie/config"

	"cloud.google.com/go/storage"
)

type storageRepo struct {
	st        *storage.BucketHandle
	appConfig *config.Config
}

//go:generate mockery --name=Storage --output=../mocks
type Storage interface {
	Upload(ctx context.Context, filename string, file multipart.File) (urlPath string, err error)
}

func New(storageClient *storage.Client, appConfig *config.Config) Storage {
	return &storageRepo{
		st:        storageClient.Bucket(appConfig.BukgetName),
		appConfig: appConfig,
	}
}
