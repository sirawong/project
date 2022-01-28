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

func TestUpdateMe(t *testing.T) {
	var err error
	mockInput := &input.UserInput{ID: "123"}
	mockOutput := &output.User{}

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/me", ctrl.UpdateMe)
		ctx.Request, err = request.MakeUpdateMeReq(ctx, mockInput)
		assert.Nil(t, err)

		ctx.Set("userid", mockInput.ID)
		userSrv.On("Update", mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.UpdateMe(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/me", ctrl.UpdateMe)
		ctx.Request, err = request.MakeUpdateMeReq(ctx, mockInput)
		assert.Nil(t, err)

		ctx.Set("userid", mockInput.ID)
		userSrv.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		ctrl.UpdateMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/me", ctrl.UpdateMe)
		ctx.Request, err = request.MakeUpdateMeReqInvalidJSON(mockInput)
		assert.Nil(t, err)

		ctrl.UpdateMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.PATCH("/users/me", ctrl.UpdateMe)
		ctx.Request, err = request.MakeUpdateMeReqInvalidJSON(mockInput)
		assert.Nil(t, err)

		ctx.Set("userid", mockInput.ID)
		ctrl.UpdateMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
