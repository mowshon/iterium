package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	assert.Exactly(t, []string{"A", "A", "A"}, Slice(Repeat("A", 3)))
}

func TestRepeatEmpty(t *testing.T) {
	assert.Exactly(t, []string{}, Slice(Repeat("A", 0)))
}

func TestRepeatInfinite(t *testing.T) {
	assert.Exactly(t, []string{"A", "A", "A", "A"}, SliceN(Repeat("A", -1), 4))
}
