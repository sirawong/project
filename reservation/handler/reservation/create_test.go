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

func TestCreate(t *testing.T) {
	var err error
	mockInput := &input.ReservationInput{}
	reservation := &output.Reservation{}

	t.Run("Success", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		reservationSrv.On("Create", mock.Anything, mockInput).Return(reservation, "", nil)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusCreated, ctx.Writer.Status())
	})

	t.Run("Error: json invalid", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeCreateReqInvalidJSON()
		assert.Nil(t, err)

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.POST("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeCreateReq(mockInput)
		assert.Nil(t, err)

		reservationSrv.On("Create", mock.Anything, mockInput).Return(nil, "", errors.New("error"))

		ctrl.Create(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
