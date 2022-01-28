package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"showtime/service/showtime/implement"
	"showtime/service/showtime/input"

	mocksRepo "showtime/repository/mocks"
	mocksUUID "showtime/utils/mocks"
)

func TestCreateShowtime(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	inputData := &input.CreateInput{
		StartAt: "18:00",
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		uuid.On("Generate").Return("1")
		repo.On("Create", ctx, mock.Anything).Return("1", nil)

		service := implement.New(repo, uuid)
		items, err := service.Create(ctx, &input.CreateInput{StartAt: "18:00"})
		assert.Nil(t, err)
		assert.Equal(t, inputData.StartAt, items.StartAt)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		uuid.On("Generate").Return("1")
		repo.On("Create", ctx, mock.Anything).Return("1", errors.New("error"))

		service := implement.New(repo, uuid)
		_, err := service.Create(ctx, &input.CreateInput{})
		assert.NotNil(t, err)
	})
}
