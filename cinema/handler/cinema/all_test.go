package cinema_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"cinema/entities"
	controller "cinema/handler/cinema"
	"cinema/handler/cinema/request"
	mockscinema "cinema/service/cinema/mocks"
	"cinema/service/cinema/output"
)

func TestAll(t *testing.T) {
	var err error
	mockInput := &entities.PageOption{}
	cinema := []*output.Cinema{}

	t.Run("Success", func(t *testing.T) {
		cinemaSrv := &mockscinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		cinemaSrv.On("List", mock.Anything, mockInput).Return(cinema, nil)

		ctrl.All(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		cinemaSrv := &mockscinema.CinemaService{}
		ctrl := controller.New(cinemaSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/cinemas", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		cinemaSrv.On("List", mock.Anything, mockInput).Return(nil, errors.New("error"))

		ctrl.All(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
