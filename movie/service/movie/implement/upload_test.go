package implement_test

import (
	"bytes"
	"context"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"movie/config"
	"movie/service/movie/implement"
	"movie/service/movie/input"

	mocksRepo "movie/repository/mocks"
	mocksUUID "movie/utils/mocks"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	uuid := &mocksUUID.UUID{}
	mockBody := &bytes.Buffer{}

	mockInput := &input.MovieInput{ID: "1"}
	appConfig := config.Get()

	fileHeader := multipart.FileHeader{Filename: "filename"}
	file, _ := fileHeader.Open()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		writer := multipart.NewWriter(mockBody)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		storage.On("Upload", ctx, fileHeader.Filename, file).Return("", nil)
		repo.On("Update", ctx, mock.Anything, mock.Anything).Return(nil)
		repo.On("Read", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		service := implement.New(repo, uuid, appConfig, storage)
		_, err := service.Upload(ctx, mockInput, fileHeader.Filename, file)
		assert.Nil(t, err)
	})

	t.Run("Error: upload", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		writer := multipart.NewWriter(mockBody)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		storage.On("Upload", ctx, fileHeader.Filename, file).Return("", errors.New("error"))

		service := implement.New(repo, uuid, appConfig, storage)
		_, err := service.Upload(ctx, mockInput, fileHeader.Filename, file)
		assert.NotNil(t, err)
	})

}
