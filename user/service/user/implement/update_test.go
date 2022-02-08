package implement_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"user/config"
	"user/entities"
	"user/service/user/implement"
	"user/service/user/input"

	mocksRepo "user/repository/mocks"
	mocksAuth "user/service/auth/mocks"
	mocksUUID "user/utils/mocks"
)

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	uuid := &mocksUUID.UUID{}
	auth := &mocksAuth.Service{}

	mockInput := &input.UserInput{
		ID:       "1",
		Name:     "dev",
		Password: "test",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	appConfig := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("Read", ctx, filters, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args[2].(*entities.User)
			arg.ID = mockInput.ID
		})
		repo.On("Update", ctx, filters, mock.Anything).Return(nil)

		service := implement.New(repo, auth, uuid, appConfig, storage)
		item, err := service.Update(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, item.ID)
	})

	t.Run("Error: Update", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("Read", ctx, filters, mock.Anything).Return(nil)
		repo.On("Update", ctx, filters, mock.Anything).Return(errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig, storage)
		_, err := service.Update(ctx, mockInput)
		assert.NotNil(t, err)
	})

	t.Run("GenPassword", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		mockInput.Password = "1"
		storage := &mocksRepo.Storage{}

		repo.On("Read", ctx, filters, mock.Anything).Return(nil)
		repo.On("Update", ctx, filters, mock.Anything).Return(nil)

		service := implement.New(repo, auth, uuid, appConfig, storage)
		_, err := service.Update(ctx, mockInput)
		assert.Nil(t, err)
	})
}
