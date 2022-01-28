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

func TestDeleteCinema(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.MovieInput{
		ID: "1",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Delete", ctx, filters).Return(nil)

		service := implement.New(repo, uuid, appConfig)
		err := service.Delete(ctx, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Delete", ctx, filters).Return(errors.New("error"))

		service := implement.New(repo, uuid, appConfig)
		err := service.Delete(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
