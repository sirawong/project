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
	"reservation/service/reservation/output"
)

func TestRead(t *testing.T) {
	var err error
	mockInput := &input.ReservationInput{ID: "123"}
	mockOutput := &output.Reservation{}

	t.Run("Success", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)
		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/reservations/:id", ctrl.Read)
		ctx.Request, err = request.MakeReadReq(ctx, mockInput)
		assert.Nil(t, err)

		reservationSrv.On("Read", mock.Anything, mock.Anything).Return(mockOutput, nil)

		ctrl.Read(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error:service", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/reservations/:id", ctrl.Read)
		ctx.Request, err = request.MakeReadReq(ctx, mockInput)
		assert.Nil(t, err)

		reservationSrv.On("Read", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		ctrl.Read(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
