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

func TestRead(t *testing.T) {
	var err error
	mockInput := &input.MovieInput{ID: "123"}
	mockOutput := &output.Movie{}

	t.Run("Success", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/movies/:id", ctrl.Read)
		ctx.Request, err = request.MakeReadReq(ctx, mockInput)
		assert.Nil(t, err)

		movieSrv.On("Read", mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Read(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		movieSrv := &mocksMovie.Service{}
		ctrl := controller.New(movieSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/movies/:id", ctrl.Read)
		ctx.Request, err = request.MakeReadReq(ctx, mockInput)
		assert.Nil(t, err)

		movieSrv.On("Read", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		ctrl.Read(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
