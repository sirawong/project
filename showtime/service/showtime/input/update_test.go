package input_test

import (
	"showtime/service/showtime/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateParseToEntities(t *testing.T) {
	mockInput := &input.UpdateInput{
		ID:      "1",
		StartAt: "18:00",
	}
	t.Run("Success", func(t *testing.T) {
		ent := mockInput.ParseToEntities()
		assert.Equal(t, mockInput.ID, ent.ID)
		assert.Equal(t, mockInput.StartAt, ent.StartAt)
	})
}
