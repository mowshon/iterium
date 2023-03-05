package iterium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator_SetInfinite(t *testing.T) {
	iter := New(1, 2, 3)
	assert.Exactly(t, false, iter.IsInfinite())

	iter.SetInfinite(true)
	assert.Exactly(t, true, iter.IsInfinite())
}
