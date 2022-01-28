package output_test

import (
	"cinema/entities"
	"cinema/service/cinema/output"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToOutput(t *testing.T) {
	mockOutput := &entities.Cinema{
		ID:   "1",
		Name: "dev",
	}
	t.Run("Success", func(t *testing.T) {
		cinema := output.ParseToOutput(mockOutput)
		assert.Equal(t, mockOutput.ID, cinema.ID)
		assert.Equal(t, mockOutput.Name, cinema.Name)
	})
}
