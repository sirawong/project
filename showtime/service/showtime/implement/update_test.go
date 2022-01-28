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

func TestUpdateShowtime(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.UpdateInput{
		ID:      "1",
		StartAt: "18:00",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(nil)

		service := implement.New(repo, uuid)
		showtime, err := service.Update(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, showtime.ID)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Update", ctx, filters, mockInput.ParseToEntities()).Return(errors.New("error"))

		service := implement.New(repo, uuid)
		_, err := service.Update(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
