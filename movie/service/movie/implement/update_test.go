package implement_test

import (
	"context"
	"errors"
	"fmt"
	"movie/config"
	mocksRepo "movie/repository/mocks"
	"movie/service/movie/implement"
	"movie/service/movie/input"
	mocksUUID "movie/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.MovieInput{
		ID:    "1",
		Title: "dev",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	ctx := context.Background()
	appConfig := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(nil)

		service := implement.New(repo, uuid, appConfig)
		cinema, err := service.Update(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, cinema.ID)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(errors.New("error"))

		service := implement.New(repo, uuid, appConfig)
		_, err := service.Update(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
