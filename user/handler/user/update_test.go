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

func TestUpdate(t *testing.T) {
	var err error
	mockInput := &input.UserInput{ID: "123"}
	mockOutput := &output.User{}

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/:id", ctrl.Update)
		ctx.Request, err = request.MakeUpdateReq(ctx, mockInput)
		assert.Nil(t, err)

		ctx.Set("userid", "456")
		userSrv.On("Update", mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Update(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/:id", ctrl.Update)
		ctx.Request, err = request.MakeUpdateReq(ctx, mockInput)
		assert.Nil(t, err)

		ctx.Set("userid", "456")
		userSrv.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		ctrl.Update(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/:id", ctrl.Update)
		ctx.Request, err = request.MakeUpdateReqInvalidJSON(mockInput)
		assert.Nil(t, err)
		
		ctx.Set("userid", "456")

		ctrl.Update(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: user admin", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/:id", ctrl.Update)
		ctx.Request, err = request.MakeUpdateReq(ctx, mockInput)
		assert.Nil(t, err)
		
		ctx.Set("userid", mockInput.ID)

		ctrl.Update(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
