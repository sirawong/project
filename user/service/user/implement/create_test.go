package implement_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"user/config"
	"user/service/user/implement"
	"user/service/user/input"

	mocksRepo "user/repository/mocks"
	mocksAuth "user/service/auth/mocks"
	mocksUUID "user/utils/mocks"
)

func TestCreate(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	auth := &mocksAuth.Service{}
	inputData := &input.UserInput{}
	name := "dev"
	inputData.Name = name
	token := "token"

	ctx := context.Background()
	appConfig := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		uuid.On("Generate").Return("id")
		repo.On("Create", ctx, mock.Anything).Return("id", nil)
		auth.On("GenAuth", ctx, mock.Anything).Return(&token, nil)

		service := implement.New(repo, auth, uuid, appConfig)
		items, _, err := service.Create(ctx, &input.UserInput{Name: name})
		assert.Nil(t, err)
		assert.Equal(t, inputData.Name, items.Name)
		assert.Equal(t, token, token)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		uuid.On("Generate").Return("id")
		repo.On("Create", ctx, mock.Anything).Return("", errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig)
		_, _, err := service.Create(ctx, &input.UserInput{})
		assert.NotNil(t, err)
	})
}
