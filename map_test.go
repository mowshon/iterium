package iterium

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapInteger(t *testing.T) {
	values := Map(New(1, 2, 3, 4, 5), func(value int) int {
		return value * 2
	})

	assert.Exactly(t, []int{2, 4, 6, 8, 10}, Slice(values))
}

func TestMapString(t *testing.T) {
	values := Map(New("a", "b", "c", "d"), strings.ToUpper)

	assert.Exactly(t, []string{"A", "B", "C", "D"}, Slice(values))
}

func TestRangeMapFilter(t *testing.T) {
	values := Filter(
		Map(Range(0, 10), func(value int) int {
			return value + 1
		}),
		func(value int) bool {
			return value%2 == 0
		},
	)

	assert.Exactly(t, []int{2, 4, 6, 8, 10}, Slice(values))
}
