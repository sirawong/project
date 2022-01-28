package input_test

import (
	"reservation/service/reservation/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateParseToEntities(t *testing.T) {
	mockInput := &input.ReservationInput{
		ID:      "1",
		StartAt: "18:00",
	}
	t.Run("Success", func(t *testing.T) {
		ent := mockInput.ParseToEntities()
		assert.Equal(t, mockInput.ID, ent.ID)
		assert.Equal(t, mockInput.StartAt, ent.StartAt)
	})
}
