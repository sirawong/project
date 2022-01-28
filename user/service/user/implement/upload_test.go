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

	"user/config"
	"user/service/user/implement"
	"user/service/user/input"

	mocksRepo "user/repository/mocks"
	mocksAuth "user/service/auth/mocks"
	mockHttp "user/service/user/mocks"
	mocksUUID "user/utils/mocks"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()
	uuid := &mocksUUID.UUID{}
	auth := &mocksAuth.Service{}
	mockBody := &bytes.Buffer{}

	mockInput := &input.UserInput{ID: "1"}
	appConfig := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
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

		service := implement.New(repo, auth, uuid, appConfig)
		_, err := service.Upload(ctx, httpClient, mockBody, writer, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error: http client", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		httpClient := &mockHttp.HttpClienter{}

		writer := multipart.NewWriter(mockBody)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		httpClient.On("Do", mock.Anything).Return(nil, errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig)
		_, err := service.Upload(ctx, httpClient, mockBody, writer, mockInput)
		assert.NotNil(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
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

		service := implement.New(repo, auth, uuid, appConfig)
		_, err := service.Upload(ctx, httpClient, mockBody, writer, mockInput)
		assert.NotNil(t, err)
	})

}
