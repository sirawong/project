package implement_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"cinema/config"
	mocksRepo "cinema/repository/mocks"
	"cinema/service/cinema/implement"
	"cinema/service/cinema/input"
	mocksUUID "cinema/utils/mocks"
)

func TestDeleteCinema(t *testing.T) {
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

		cinemaRepo.On("Delete", ctx, filters).Return(nil)

		service := implement.New(cinemaRepo, uuid, config, storage)
		err := service.Delete(ctx, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		cinemaRepo := &mocksRepo.CinemaRepository{}
		storage := &mocksRepo.Storage{}

		cinemaRepo.On("Delete", ctx, filters).Return(errors.New("error"))

		service := implement.New(cinemaRepo, uuid, config, storage)
		err := service.Delete(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
