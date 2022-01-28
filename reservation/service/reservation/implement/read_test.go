package implement_test

import (
	"context"
	"errors"
	"fmt"
	"reservation/config"
	"reservation/entities"
	mocksRepo "reservation/repository/mocks"
	mocksGRPC "reservation/service/grpcClient/mocks"
	"reservation/service/reservation/implement"
	"reservation/service/reservation/input"
	mocksUUID "reservation/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReadCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	grpc := &mocksGRPC.Service{}
	mockInput := &input.ReservationInput{
		ID: "1",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Read", ctx, filters, &entities.Reservation{}).Return(nil).Run(func(args mock.Arguments) {
			arg := args[2].(*entities.Reservation)
			arg.ID = mockInput.ID
		})

		service := implement.New(repo, uuid, grpc, appConfig)
		cinema, err := service.Read(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, cinema.ID)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Read", ctx, filters, &entities.Reservation{}).Return(errors.New("error"))

		service := implement.New(repo, uuid, grpc, appConfig)
		_, err := service.Read(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
