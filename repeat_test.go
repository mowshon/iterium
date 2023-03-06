package iterium

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

func TestRepeatInfinite(t *testing.T) {
	repeat := Repeat("A", -1)
	assert.Exactly(t, true, repeat.IsInfinite())

	var data []string
	for i := 0; i <= 3; i++ {
		if next, err := repeat.Next(); assert.Nil(t, err) {
			data = append(data, next)
		}
	}

	assert.Exactly(t, []string{"A", "A", "A", "A"}, data)
}