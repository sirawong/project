package implement_test

import (
	"context"
	"errors"
	"testing"

	"cinema/config"
	"cinema/entities"
	mocksRepo "cinema/repository/mocks"
	"cinema/service/cinema/implement"
	mocksUUID "cinema/utils/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	opt := &entities.PageOption{
		Page:    0,
		PerPage: 0,
	}

	mockItmes := make([]interface{}, 1)
	for i := 0; i < 1; i++ {
		mockItmes[i] = &entities.Cinema{}
	}

	config := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}

		cinemaRepo.On("List", ctx, opt, &entities.Cinema{}).Return(1, mockItmes, nil)

		service := implement.New(cinemaRepo, uuid, config)
		items, err := service.List(ctx, opt)
		assert.Nil(t, err)
		assert.Equal(t, len(mockItmes), len(items))
	})

	t.Run("Error", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}

		cinemaRepo.On("List", ctx, opt, &entities.Cinema{}).Return(1, mockItmes, errors.New("error"))

		service := implement.New(cinemaRepo, uuid, config)
		_, err := service.List(ctx, opt)
		assert.NotNil(t, err)
	})
}
