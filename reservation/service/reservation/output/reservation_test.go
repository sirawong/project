package output_test

import (
	"reservation/entities"
	"reservation/service/reservation/output"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToOutput(t *testing.T) {
	mockOutput := &entities.Reservation{
		ID:      "1",
		StartAt: "18:00",
	}
	t.Run("Success", func(t *testing.T) {
		reservation := output.ParseToOutput(mockOutput)
		assert.Equal(t, mockOutput.ID, reservation.ID)
		assert.Equal(t, mockOutput.StartAt, reservation.StartAt)
	})
}
