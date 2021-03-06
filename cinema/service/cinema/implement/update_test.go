package implement_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"cinema/config"
	"cinema/service/cinema/implement"
	"cinema/service/cinema/input"

	mocksRepo "cinema/repository/mocks"
	mocksUUID "cinema/utils/mocks"
)

func TestUpdateCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.CinemaInput{
		ID:   "1",
		Name: "dev",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	config := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		cinemaRepo.On("Read", ctx, filters, mock.Anything).Return(nil)
		cinemaRepo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(nil)

		service := implement.New(cinemaRepo, uuid, config, storage)
		cinema, err := service.Update(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, cinema.ID)
	})

	t.Run("Error", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		cinemaRepo.On("Read", ctx, filters, mock.Anything).Return(nil)
		cinemaRepo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(errors.New("error"))

		service := implement.New(cinemaRepo, uuid, config, storage)
		_, err := service.Update(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
