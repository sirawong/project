package reservation_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	controller "reservation/handler/reservation"
	"reservation/handler/reservation/request"
	"reservation/service/reservation/input"
	mocksReservation "reservation/service/reservation/mocks"
)

func TestDelete(t *testing.T) {
	var err error
	mockInput := &input.ReservationInput{ID: "123"}

	t.Run("Success", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		reservationSrv.On("Delete", mock.Anything, mockInput).Return(nil)

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.DELETE("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeDeleteReq(ctx, mockInput)
		assert.Nil(t, err)

		reservationSrv.On("Delete", mock.Anything, mockInput).Return(errors.New("error"))

		ctrl.Delete(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
