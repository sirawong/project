package movie_test

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	controller "movie/handler/movie"
	"movie/handler/movie/request"
	"movie/service/movie/input"
	mocksMovie "movie/service/movie/mocks"
	"movie/service/movie/output"
)

func TestUpload(t *testing.T) {
	var err error
	mockInput := &input.MovieInput{ID: "123"}
	mockOutput := &output.Movie{}

	body := new(bytes.Buffer)

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksMovie.Service{}
		ctrl := controller.New(userSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/photo/:id", ctrl.Upload)

		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		ctx.Request, err = request.MakeUploadReq(mockInput.ID, body, writer)
		assert.Nil(t, err)

		userSrv.On("Upload", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Upload(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error: formfile", func(t *testing.T) {
		userSrv := &mocksMovie.Service{}
		ctrl := controller.New(userSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/photo/:id", ctrl.Upload)

		writer := multipart.NewWriter(body)

		ctx.Request, err = request.MakeUploadReq(mockInput.ID, &bytes.Buffer{}, writer)
		assert.Nil(t, err)

		userSrv.On("Upload", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Upload(ctx)

		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		userSrv := &mocksMovie.Service{}
		ctrl := controller.New(userSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/photo/:id", ctrl.Upload)

		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		ctx.Request, err = request.MakeUploadReq(mockInput.ID, body, writer)
		assert.Nil(t, err)

		userSrv.On("Upload", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("erro"))

		ctrl.Upload(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

}
