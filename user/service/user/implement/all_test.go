package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"user/config"
	"user/entities"
	"user/service/user/implement"

	mocksRepo "user/repository/mocks"
	mocksAuth "user/service/auth/mocks"
	mocksUUID "user/utils/mocks"
)

func TestGetAll(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	auth := &mocksAuth.Service{}
	opt := &entities.PageOption{
		Page:    0,
		PerPage: 0,
	}

	mockItmes := make([]interface{}, 1)
	for i := 0; i < 1; i++ {
		mockItmes[i] = &entities.User{}
	}

	appConfig := config.Get()
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("List", ctx, opt, &entities.User{}).Return(1, mockItmes, nil)

		service := implement.New(repo, auth, uuid, appConfig)
		items, err := service.All(ctx)
		assert.Nil(t, err)
		assert.Equal(t, len(mockItmes), len(items))
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("List", ctx, opt, &entities.User{}).Return(1, mockItmes, errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig)
		_, err := service.All(ctx)
		assert.NotNil(t, err)
	})
}
