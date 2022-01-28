package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"user/entities"
	controller "user/handler/user"
	"user/handler/user/request"
	mocksuser "user/service/user/mocks"
	"user/service/user/output"
)

func TestAll(t *testing.T) {
	var err error
	mockInput := &entities.PageOption{}
	user := []*output.User{}

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksuser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/users", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		userSrv.On("List", mock.Anything, mockInput).Return(user, nil)

		ctrl.All(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		userSrv := &mocksuser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/users", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		userSrv.On("List", mock.Anything, mockInput).Return(nil, errors.New("error"))

		ctrl.All(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
