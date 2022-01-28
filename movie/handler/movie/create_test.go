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
	"movie/service/movie/output"
)

func TestCreate(t *testing.T) {
	var err error
	mockInput := &input.MovieInput{}
	movie := &output.Movie{}

	t.Run("Success", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/movies", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		movieSrv.On("Create", mock.Anything, mockInput).Return(movie, nil)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusCreated, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/movies", ctrl.Create)
		ctx.Request, err = request.MakeCreateReqInvalidJSON()
		assert.Nil(t, err)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/movies", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		movieSrv.On("Create", mock.Anything, mockInput).Return(nil, errors.New("error"))

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
