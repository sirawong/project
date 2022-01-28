package movie_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"movie/entities"
	controller "movie/handler/movie"
	"movie/handler/movie/request"
	mocksMovie "movie/service/movie/mocks"
	"movie/service/movie/output"
)

func TestAll(t *testing.T) {
	var err error
	mockInput := &entities.PageOption{}
	movie := []*output.Movie{}

	t.Run("Success", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/movies", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		movieSrv.On("List", mock.Anything, mockInput).Return(movie, nil)

		ctrl.All(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/movies", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		movieSrv.On("List", mock.Anything, mockInput).Return(nil, errors.New("error"))

		ctrl.All(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
