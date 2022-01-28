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

func TestDeleteCinema(t *testing.T) {
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

		repo.On("Delete", ctx, filters).Return(nil)

		service := implement.New(repo, uuid, grpc, appConfig)
		err := service.Delete(ctx, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Delete", ctx, filters).Return(errors.New("error"))

		service := implement.New(repo, uuid, grpc, appConfig)
		err := service.Delete(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
