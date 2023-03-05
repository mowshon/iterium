package iterium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func lessThen5(x int) bool {
	return x < 5
}

func TestFilterFalse(t *testing.T) {
	filterfalse := FilterFalse(New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), lessThen5)
	if slice, err := filterfalse.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{5, 6, 7, 8, 9}, slice)
	}
}
