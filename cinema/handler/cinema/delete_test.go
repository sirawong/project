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
)

func TestDelete(t *testing.T) {
	var err error
	mockInput := &input.CinemaInput{ID: "123"}

	t.Run("Success", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		cinemaSrv.On("Delete", mock.Anything, mockInput).Return(nil)

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		cinemaSrv := &mocksCinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		cinemaSrv.On("Delete", mock.Anything, mockInput).Return(errors.New("error"))

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
