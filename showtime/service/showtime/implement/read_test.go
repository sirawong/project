package implement_test

import (
	"context"
	"errors"
	"fmt"
	"showtime/entities"
	mocksRepo "showtime/repository/mocks"
	"showtime/service/showtime/implement"
	"showtime/service/showtime/input"
	mocksUUID "showtime/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReadShowtime(t *testing.T) {
	uuid := &mocksUUID.UUID{}
	mockInput := &input.ReadInput{
		ID: "1",
	}

	filters := []string{
		fmt.Sprintf("_id:eq:%v", mockInput.ID),
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Read", ctx, filters, &entities.ShowTime{}).Return(nil).Run(func(args mock.Arguments) {
			arg := args[2].(*entities.ShowTime)
			arg.ID = mockInput.ID
		})

		service := implement.New(repo, uuid)
		showtime, err := service.Read(ctx, mockInput)
		assert.Nil(t, err)
		assert.Equal(t, mockInput.ID, showtime.ID)
	})

	t.Run("Error", func(t *testing.T) {
		repo := &mocksRepo.Repository{}

		repo.On("Read", ctx, filters, &entities.ShowTime{}).Return(errors.New("error"))

		service := implement.New(repo, uuid)
		_, err := service.Read(ctx, mockInput)
		assert.NotNil(t, err)
	})
}
