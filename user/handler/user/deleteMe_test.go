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
)

func TestDeleteMe(t *testing.T) {
	var err error
	mockInput := &input.UserInput{ID: "123"}
	// mockOutput := &output.User{}

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/users/me", ctrl.DeleteMe)
		ctx.Request, err = request.MakeDeleteMeReq()
		assert.Nil(t, err)

		mockInput.Role = "superadmin"
		ctx.Set("userid", mockInput.ID)
		ctx.Set("role", mockInput.Role)
		userSrv.On("Delete", mock.Anything, mock.Anything).Return(nil)

		ctrl.DeleteMe(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/users/me", ctrl.DeleteMe)
		ctx.Request, err = request.MakeDeleteMeReq()
		assert.Nil(t, err)

		mockInput.Role = "superadmin"
		ctx.Set("userid", mockInput.ID)
		ctx.Set("role", mockInput.Role)
		userSrv.On("Delete", mock.Anything, mock.Anything).Return(errors.New("error"))

		ctrl.DeleteMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error:get userid", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/users/me", ctrl.DeleteMe)
		ctx.Request, err = request.MakeDeleteMeReq()
		assert.Nil(t, err)

		mockInput.Role = "superadmin"
		ctx.Set("role", mockInput.Role)
		ctrl.DeleteMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error:get role", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/users/me", ctrl.GetMe)
		ctx.Request, err = request.MakeDeleteMeReq()
		assert.Nil(t, err)

		ctrl.DeleteMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error:role not superadmin", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/users/me", ctrl.GetMe)
		ctx.Request, err = request.MakeDeleteMeReq()
		assert.Nil(t, err)

		mockInput.Role = "guest"
		ctx.Set("role", mockInput.Role)
		ctrl.DeleteMe(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
