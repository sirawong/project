package implement_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"cinema/config"
	"cinema/service/cinema/implement"
	"cinema/service/cinema/input"

	mocksRepo "cinema/repository/mocks"
	mockHttp "cinema/service/cinema/mocks"
	mocksUUID "cinema/utils/mocks"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	uuid := &mocksUUID.UUID{}
	mockBody := &bytes.Buffer{}

	mockInput := &input.CinemaInput{ID: "1"}
	config := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.CinemaRepository{}
		httpClient := &mockHttp.HttpClienter{}

		writer := multipart.NewWriter(mockBody)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		st := strings.NewReader("string")
		stringReadCloser := io.NopCloser(st)

		resp := &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

		httpClient.On("Do", mock.Anything).Return(resp, nil)
		repo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		service := implement.New(repo, uuid, config)
		_, err := service.Upload(ctx, httpClient, mockBody, writer, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error: http client", func(t *testing.T) {
		repo := &mocksRepo.CinemaRepository{}
		httpClient := &mockHttp.HttpClienter{}

		writer := multipart.NewWriter(mockBody)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		httpClient.On("Do", mock.Anything).Return(nil, errors.New("error"))

		service := implement.New(repo, uuid, config)
		_, err := service.Upload(ctx, httpClient, mockBody, writer, mockInput)
		assert.NotNil(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.CinemaRepository{}
		httpClient := &mockHttp.HttpClienter{}

		writer := multipart.NewWriter(mockBody)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		st := strings.NewReader("string")
		stringReadCloser := io.NopCloser(st)

		resp := &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

		httpClient.On("Do", mock.Anything).Return(resp, nil)
		repo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

		service := implement.New(repo, uuid, config)
		_, err := service.Upload(ctx, httpClient, mockBody, writer, mockInput)
		assert.NotNil(t, err)
	})

}
