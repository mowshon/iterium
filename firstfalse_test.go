package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstFalse(t *testing.T) {
	value, ok := FirstFalse(Range(10, 0, -1), func(value int) bool {
		return value >= 2
	})

	assert.True(t, ok)
	assert.Exactly(t, 1, value)
}

func TestFirstFalseEmpty(t *testing.T) {
	_, ok := FirstFalse(Range(10), func(value int) bool {
		return value < 100
	})

	assert.False(t, ok)
}
