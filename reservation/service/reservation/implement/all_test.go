package implement_test

import (
	"context"
	"errors"
	"testing"

	"reservation/config"
	"reservation/entities"
	"reservation/service/reservation/implement"

	mocksRepo "reservation/repository/mocks"
	mocksGRPC "reservation/service/grpcClient/mocks"
	mocksUUID "reservation/utils/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	grpc := &mocksGRPC.Service{}
	opt := &entities.PageOption{
		Page:    0,
		PerPage: 0,
	}

	mockItmes := make([]interface{}, 1)
	for i := 0; i < 1; i++ {
		mockItmes[i] = &entities.Reservation{}
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("List", ctx, opt, &entities.Reservation{}).Return(1, mockItmes, nil)

		service := implement.New(repo, uuid, grpc, appConfig)
		items, err := service.List(ctx, opt)
		assert.Nil(t, err)
		assert.Equal(t, len(mockItmes), len(items))
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("List", ctx, opt, &entities.Reservation{}).Return(1, mockItmes, errors.New("error"))

		service := implement.New(repo, uuid, grpc, appConfig)
		_, err := service.List(ctx, opt)
		assert.NotNil(t, err)
	})
}
