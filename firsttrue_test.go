package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstTrue(t *testing.T) {
	value, ok := FirstTrue(Range(10), func(value int) bool {
		return value == 5
	})

	assert.True(t, ok)
	assert.Exactly(t, 5, value)
}

func TestFirstTrueEmpty(t *testing.T) {
	_, ok := FirstTrue(Range(10), func(value int) bool {
		return value > 100
	})

	assert.False(t, ok)
}
