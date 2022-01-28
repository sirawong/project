package implement_test

import (
	"context"
	"errors"
	"fmt"
	mocksRepo "showtime/repository/mocks"
	"showtime/service/showtime/implement"
	"showtime/service/showtime/input"
	mocksUUID "showtime/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteShowtime(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.DeleteInput{
		ID: "1",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Delete", ctx, filters).Return(nil)

		service := implement.New(repo, uuid)
		err := service.Delete(ctx, mockInput)
		assert.Nil(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Delete", ctx, filters).Return(errors.New("error"))

		service := implement.New(repo, uuid)
		err := service.Delete(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
