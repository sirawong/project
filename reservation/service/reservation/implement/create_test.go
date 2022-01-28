package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"reservation/config"
	"reservation/service/reservation/implement"
	"reservation/service/reservation/input"

	mocksRepo "reservation/repository/mocks"
	mocksGRPC "reservation/service/grpcClient/mocks"
	mocksUtils "reservation/utils/mocks"
)

func TestCreateMovie(t *testing.T) {
	uuid := &mocksUtils.UUID{}
	grpc := &mocksGRPC.Service{}
	inputData := &input.ReservationInput{
		StartAt: "18:00",
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		uuid.On("Generate").Return("1")
		repo.On("Create", ctx, mock.Anything).Return("1", nil)

		service := implement.New(repo, uuid, grpc, appConfig)
		items, _, err := service.Create(ctx, &input.ReservationInput{StartAt: "18:00"})
		assert.Nil(t, err)
		assert.Equal(t, inputData.StartAt, items.StartAt)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		uuid.On("Generate").Return("1")
		repo.On("Create", ctx, mock.Anything).Return("1", errors.New("error"))

		service := implement.New(repo, uuid, grpc, appConfig)
		_, _, err := service.Create(ctx, &input.ReservationInput{})
		assert.NotNil(t, err)
	})
}
