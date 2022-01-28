package implement_test

import (
	"context"
	"errors"
	"showtime/entities"
	mocksRepo "showtime/repository/mocks"
	"showtime/service/showtime/implement"
	mocksUUID "showtime/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllShowtime(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	opt := &entities.PageOption{
		Page:    0,
		PerPage: 0,
	}

	mockItmes := make([]interface{}, 1)
	for i := 0; i < 1; i++ {
		mockItmes[i] = &entities.ShowTime{}
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("List", ctx, opt, &entities.ShowTime{}).Return(1, mockItmes, nil)

		service := implement.New(repo, uuid)
		items, err := service.All(ctx)
		assert.Nil(t, err)
		assert.Equal(t, len(mockItmes), len(items))
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("List", ctx, opt, &entities.ShowTime{}).Return(1, mockItmes, errors.New("error"))

		service := implement.New(repo, uuid)
		_, err := service.All(ctx)
		assert.NotNil(t, err)
	})
}
