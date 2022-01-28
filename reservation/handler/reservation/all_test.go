package reservation_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"reservation/entities"
	controller "reservation/handler/reservation"
	"reservation/handler/reservation/request"
	mocksReservation "reservation/service/reservation/mocks"
	"reservation/service/reservation/output"
)

func TestAll(t *testing.T) {
	var err error
	mockInput := &entities.PageOption{}
	reservation := []*output.Reservation{}

	t.Run("Success", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		reservationSrv.On("List", mock.Anything, mockInput).Return(reservation, nil)

		ctrl.All(ctx)

		assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	})

	t.Run("Error: service", func(t *testing.T) {
		reservationSrv := &mocksReservation.Service{}
		ctrl := controller.New(reservationSrv)

		ctx, router := gin.CreateTestContext(httptest.NewRecorder())
		router.GET("/reservations", ctrl.Create)
		ctx.Request, err = request.MakeAllReq(mockInput)
		assert.Nil(t, err)

		reservationSrv.On("List", mock.Anything, mockInput).Return(nil, errors.New("error"))

		ctrl.All(ctx)

		assert.Equal(t, http.StatusInternalServerError, ctx.Writer.Status())
	})
}
