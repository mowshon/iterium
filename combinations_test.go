package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	values := joinStrings(Combinations([]string{"A", "B", "C", "D"}, 2))

	assert.Exactly(t, []string{"AB", "AC", "AD", "BC", "BD", "CD"}, values)
}

func TestCombinationsEdges(t *testing.T) {
	assert.Exactly(t, [][]string{{}}, Slice(Combinations([]string{"A", "B"}, 0)))
	assert.Exactly(t, [][]string{}, Slice(Combinations([]string{"A", "B"}, 3)))
	assert.Exactly(t, [][]string{}, Slice(Combinations([]string{"A", "B"}, -1)))
}

func TestCombinationsCount(t *testing.T) {
	assert.Exactly(t, int64(6), CombinationsCount(4, 2))
	assert.Exactly(t, int64(171), CombinationsCount(19, 2))
	assert.Exactly(t, int64(658008), CombinationsCount(40, 5))
}
