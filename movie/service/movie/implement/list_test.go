package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"movie/config"
	"movie/entities"

	mocksRepo "movie/repository/mocks"
	"movie/service/movie/implement"
	mocksUUID "movie/utils/mocks"
)

func TestListCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	opt := &entities.PageOption{
		Page:    0,
		PerPage: 0,
	}

	mockItmes := make([]interface{}, 1)
	for i := 0; i < 1; i++ {
		mockItmes[i] = &entities.Movie{}
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("List", ctx, opt, &entities.Movie{}).Return(1, mockItmes, nil)

		service := implement.New(repo, uuid, appConfig, storage)
		items, err := service.List(ctx, opt)
		assert.Nil(t, err)
		assert.Equal(t, len(mockItmes), len(items))
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("List", ctx, opt, &entities.Movie{}).Return(1, mockItmes, errors.New("error"))

		service := implement.New(repo, uuid, appConfig, storage)
		_, err := service.List(ctx, opt)
		assert.NotNil(t, err)
	})
}
