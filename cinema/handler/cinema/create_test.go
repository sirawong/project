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

func TestCreate(t *testing.T) {
	var err error
	mockInput := &input.CinemaInput{}
	cinema := &output.Cinema{}

	t.Run("Success", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		cinemaSrv.On("Create", mock.Anything, mockInput).Return(cinema, nil)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusCreated, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeCreateReqInvalidJSON()
		assert.Nil(t, err)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		cinemaSrv.On("Create", mock.Anything, mockInput).Return(nil, errors.New("error"))

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
