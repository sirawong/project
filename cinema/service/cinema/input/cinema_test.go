package input_test

import (
	"cinema/service/cinema/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateParseToEntities(t *testing.T) {
	mockInput := &input.CinemaInput{
		ID:   "1",
		Name: "dev",
	}
	t.Run("Success", func(t *testing.T) {
		ent := mockInput.ParseToEntities()
		assert.Equal(t, mockInput.ID, ent.ID)
		assert.Equal(t, mockInput.Name, ent.Name)
	})
}
