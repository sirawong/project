package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"cinema/config"
	mocksRepo "cinema/repository/mocks"
	"cinema/service/cinema/implement"
	"cinema/service/cinema/input"
	mocksUUID "cinema/utils/mocks"
)

func TestCreateCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	cinemaInput := &input.CinemaInput{
		Name: "dev",
	}

	config := config.Get()

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}

		uuid.On("Generate").Return("1")
		cinemaRepo.On("Create", ctx, mock.Anything).Return("1", nil)

		service := implement.New(cinemaRepo, uuid, config)
		cinema, err := service.Create(ctx, &input.CinemaInput{Name: "dev"})
		assert.Nil(t, err)
		assert.Equal(t, cinemaInput.Name, cinema.Name)
	})

	t.Run("Error", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}
		uuid.On("Generate").Return("1")
		cinemaRepo.On("Create", ctx, mock.Anything).Return("1", errors.New("error"))

		service := implement.New(cinemaRepo, uuid, config)
		_, err := service.Create(ctx, &input.CinemaInput{})
		assert.NotNil(t, err)
	})
}
