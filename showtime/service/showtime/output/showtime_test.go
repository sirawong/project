package output_test

import (
	"showtime/entities"
	"showtime/service/showtime/output"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToOutput(t *testing.T) {
	mockOutput := &entities.ShowTime{
		ID:      "1",
		StartAt: "18:00",
	}
	t.Run("Success", func(t *testing.T) {
		movie := output.ParseToOutput(mockOutput)
		assert.Equal(t, mockOutput.ID, movie.ID)
		assert.Equal(t, mockOutput.StartAt, movie.StartAt)
	})
}
