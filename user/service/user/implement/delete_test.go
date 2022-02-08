package implement_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"user/config"
	"user/service/user/implement"
	"user/service/user/input"

	mocksRepo "user/repository/mocks"
	mocksAuth "user/service/auth/mocks"
	mocksUUID "user/utils/mocks"
)

func TestDelete(t *testing.T) {
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

		repo.On("Delete", ctx, filters).Return(nil)

		service := implement.New(repo, auth, uuid, appConfig, storage)
		err := service.Delete(ctx, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}
		storage := &mocksRepo.Storage{}

		repo.On("Delete", ctx, filters).Return(errors.New("error"))

		service := implement.New(repo, auth, uuid, appConfig, storage)
		err := service.Delete(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
