package utils_test

import (
	"showtime/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	t.Run("NewUUID Success", func(t *testing.T) {
		uuid, err := utils.NewUUID()
		assert.Nil(t, err)
		assert.NotNil(t, uuid)
	})
	t.Run("Generate Success", func(t *testing.T) {
		uuid, err := utils.NewUUID()
		assert.Nil(t, err)
		assert.NotNil(t, uuid.Generate())
	})
}
