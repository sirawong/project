package implement_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"cinema/config"
	"cinema/entities"
	mocksRepo "cinema/repository/mocks"
	"cinema/service/cinema/implement"
	"cinema/service/cinema/input"
	mocksUUID "cinema/utils/mocks"
)

func TestReadCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.CinemaInput{
		ID: "1",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	config := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		cinemaRepo.On("Read", ctx, filters, &entities.Cinema{}).Return(nil).Run(func(args mock.Arguments) {
			arg := args[2].(*entities.Cinema)
			arg.ID = mockInput.ID
		})

		service := implement.New(cinemaRepo, uuid, config, storage)
		cinema, err := service.Read(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, cinema.ID)
	})

	t.Run("Error", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		cinemaRepo.On("Read", ctx, filters, &entities.Cinema{}).Return(errors.New("error"))

		service := implement.New(cinemaRepo, uuid, config, storage)
		_, err := service.Read(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
