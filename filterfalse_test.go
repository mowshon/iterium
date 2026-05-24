package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lessThen5(value int) bool {
	return value < 5
}

func TestFilterFalse(t *testing.T) {
	values := FilterFalse(New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), lessThen5)

	assert.Exactly(t, []int{5, 6, 7, 8, 9}, Slice(values))
}
