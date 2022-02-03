package implement_test

import (
	"context"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"user/config"
	"user/service/user/implement"
	"user/service/user/input"

	mocksRepo "user/repository/mocks"
	mocksAuth "user/service/auth/mocks"
	mocksUUID "user/utils/mocks"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	uuid := &mocksUUID.UUID{}
	auth := &mocksAuth.Service{}

	mockInput := &input.UserInput{ID: "1"}
	appConfig := config.Get()

	fileHeader := multipart.FileHeader{Filename: "filename"}
	file, _ := fileHeader.Open()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		storage.On("Upload", ctx, fileHeader.Filename, file).Return("", nil)
		repo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		repo.On("Read", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		service := implement.New(repo, auth, uuid, appConfig, storage)
		_, err := service.Upload(ctx, mockInput, fileHeader.Filename, file)
		assert.Nil(t, err)
	})

	t.Run("Error: upload", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		storage.On("Upload", ctx, fileHeader.Filename, file).Return("", errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig, storage)
		_, err := service.Upload(ctx, mockInput, fileHeader.Filename, file)
		assert.NotNil(t, err)
	})

}
