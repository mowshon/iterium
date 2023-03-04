package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("A", 3)

	if slice, err := repeat.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []string{"A", "A", "A"}, slice)
	}
}
