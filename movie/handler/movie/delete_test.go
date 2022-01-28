package movie_test

import (
	"errors"
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
)

func TestDelete(t *testing.T) {
	var err error
	mockInput := &input.MovieInput{ID: "123"}

	t.Run("Success", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/movies", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		movieSrv.On("Delete", mock.Anything, mockInput).Return(nil)

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/movies", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		movieSrv.On("Delete", mock.Anything, mockInput).Return(errors.New("error"))

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
