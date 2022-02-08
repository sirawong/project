package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"movie/config"
	"movie/service/movie/implement"
	"movie/service/movie/input"

	mocksRepo "movie/repository/mocks"
	mocksUUID "movie/utils/mocks"
)

func TestCreateCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	inputData := &input.MovieInput{
		Title: "dev",
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		uuid.On("Generate").Return("1")
		repo.On("Create", ctx, mock.Anything).Return("1", nil)

		service := implement.New(repo, uuid, appConfig, storage)
		items, err := service.Create(ctx, &input.MovieInput{Title: "dev"})
		assert.Nil(t, err)
		assert.Equal(t, inputData.Title, items.Title)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}
		uuid.On("Generate").Return("1")
		repo.On("Create", ctx, mock.Anything).Return("1", errors.New("error"))

		service := implement.New(repo, uuid, appConfig, storage)
		_, err := service.Create(ctx, &input.MovieInput{})
		assert.NotNil(t, err)
	})
}
