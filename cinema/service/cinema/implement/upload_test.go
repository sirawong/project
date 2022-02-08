package implement_test

import (
	"context"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"cinema/config"
	"cinema/service/cinema/implement"
	"cinema/service/cinema/input"

	mocksRepo "cinema/repository/mocks"
	mocksUUID "cinema/utils/mocks"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	uuid := &mocksUUID.UUID{}

	mockInput := &input.CinemaInput{ID: "1"}
	config := config.Get()

	fileHeader := multipart.FileHeader{Filename: "filename"}
	file, _ := fileHeader.Open()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		storage.On("Upload", ctx, fileHeader.Filename, file).Return("", nil)
		repo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		repo.On("Read", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		service := implement.New(repo, uuid, config, storage)
		_, err := service.Upload(ctx, mockInput, fileHeader.Filename, file)
		assert.Nil(t, err)
	})

	t.Run("Error: http client", func(t *testing.T) {
		repo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		storage.On("Upload", ctx, fileHeader.Filename, file).Return("", errors.New("error"))

		service := implement.New(repo, uuid, config, storage)
		_, err := service.Upload(ctx, mockInput, fileHeader.Filename, file)
		assert.NotNil(t, err)
	})

}
