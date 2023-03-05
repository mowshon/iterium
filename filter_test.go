package iterium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func even(x int) bool {
	return x%2 == 0
}

func TestFilter(t *testing.T) {
	filter := Filter(New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), even)
	if slice, err := filter.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{2, 4, 6, 8, 10}, slice)
	}
}

func TestFilterEmpty(t *testing.T) {
	filter := Filter(New(1, 2, 3), func(x int) bool {
		return x > 100
	})

	if slice, err := filter.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{}, slice)
	}
}
