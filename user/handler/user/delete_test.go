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

func TestDelete(t *testing.T) {
	var err error
	mockInput := &input.UserInput{ID: "123"}

	t.Run("Success", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/users", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		userSrv.On("Delete", mock.Anything, mockInput).Return(nil)

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		userSrv := &mocksUser.Service{}
		ctrl := controller.New(userSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/users", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		userSrv.On("Delete", mock.Anything, mockInput).Return(errors.New("error"))

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
