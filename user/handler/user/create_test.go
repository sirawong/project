package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	controller "user/handler/user"
	"user/handler/user/request"
	"user/service/user/input"
	mocksUser "user/service/user/mocks"
	"user/service/user/output"
)

func TestCreate(t *testing.T) {
	var err error
	mockInput := &input.UserInput{}
	user := &output.User{}

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/users", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		token := "token"
		userSrv.On("Create", mock.Anything, mockInput).Return(user, &token, nil)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusCreated, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/users", ctrl.Create)
		ctx.Request, err = request.MakeCreateReqInvalidJSON()
		assert.Nil(t, err)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/users", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		userSrv.On("Create", mock.Anything, mockInput).Return(nil, nil, errors.New("error"))

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
