package cinema_test

import (
	"errors"
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

func TestRead(t *testing.T) {
	var err error
	mockInput := &input.CinemaInput{ID: "123"}
	mockOutput := &output.Cinema{}

	t.Run("Success", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/cinemas/:id", ctrl.Read)
		ctx.Request, err = request.MakeReadReq(ctx, mockInput)
		assert.Nil(t, err)

		ctx.Set("cinemaid", mockInput.ID)
		cinemaSrv.On("Read", mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Read(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/cinemas/:id", ctrl.Read)
		ctx.Request, err = request.MakeReadReq(ctx, mockInput)
		assert.Nil(t, err)

		ctx.Set("cinemaid", mockInput.ID)
		cinemaSrv.On("Read", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		ctrl.Read(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
