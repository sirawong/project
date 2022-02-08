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

func TestRead(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	auth := &mocksAuth.Service{}
	mockInput := &input.UserInput{
		ID: "1",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	ctx := context.Background()
	appConfig := config.Get()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("Read", ctx, filters, &entities.User{}).Return(nil).Run(func(args mock.Arguments) {
			arg := args[2].(*entities.User)
			arg.ID = mockInput.ID
		})

		service := implement.New(repo, auth, uuid, appConfig, storage)
		item, err := service.Read(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, item.ID)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("Read", ctx, filters, &entities.User{}).Return(errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig, storage)
		_, err := service.Read(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
