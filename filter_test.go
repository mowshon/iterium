package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func even(value int) bool {
	return value%2 == 0
}

func TestFilter(t *testing.T) {
	values := Filter(New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), even)

	assert.Exactly(t, []int{2, 4, 6, 8, 10}, Slice(values))
}

func TestFilterEmpty(t *testing.T) {
	values := Filter(New(1, 2, 3), func(value int) bool {
		return value > 100
	})

	assert.Exactly(t, []int{}, Slice(values))
}
