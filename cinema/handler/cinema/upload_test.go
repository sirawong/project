package cinema_test

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

	controller "cinema/handler/cinema"
	"cinema/handler/cinema/request"
	"cinema/service/cinema/input"
	mocksCinema "cinema/service/cinema/mocks"
	"cinema/service/cinema/output"
)

func TestUpload(t *testing.T) {
	var err error
	mockInput := &input.CinemaInput{ID: "123"}
	mockOutput := &output.Cinema{}

	body := new(bytes.Buffer)

	t.Run("Success", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/photo/:id", ctrl.Upload)

		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		ctx.Request, err = request.MakeUploadReq(mockInput.ID, body, writer)
		assert.Nil(t, err)

		cinemaSrv.On("Upload", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Upload(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error: formfile", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/photo/:id", ctrl.Upload)

		writer := multipart.NewWriter(body)

		ctx.Request, err = request.MakeUploadReq(mockInput.ID, &bytes.Buffer{}, writer)
		assert.Nil(t, err)

		cinemaSrv.On("Upload", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Upload(ctx)

		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/photo/:id", ctrl.Upload)

		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "file.png")
		part.Write([]byte(`sample`))
		writer.Close()

		ctx.Request, err = request.MakeUploadReq(mockInput.ID, body, writer)
		assert.Nil(t, err)

		cinemaSrv.On("Upload", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("erro"))

		ctrl.Upload(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

}
