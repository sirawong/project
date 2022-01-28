package implement_test

import (
	"context"
	"errors"
	"fmt"
	"reservation/config"
	mocksRepo "reservation/repository/mocks"
	mocksGRPC "reservation/service/grpcClient/mocks"
	"reservation/service/reservation/implement"
	"reservation/service/reservation/input"
	mocksUUID "reservation/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	grpc := &mocksGRPC.Service{}
	mockInput := &input.ReservationInput{
		ID:      "1",
		StartAt: "18:00",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	ctx := context.Background()
	appConfig := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(nil)

		service := implement.New(repo, uuid, grpc, appConfig)
		cinema, err := service.Update(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, cinema.ID)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(errors.New("error"))

		service := implement.New(repo, uuid, grpc, appConfig)
		_, err := service.Update(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
