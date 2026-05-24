package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	values := joinStrings(Permutations([]string{"A", "B", "C", "D"}, 2))

	assert.Exactly(t, []string{
		"AB", "AC", "AD", "BA", "BC", "BD",
		"CA", "CB", "CD", "DA", "DB", "DC",
	}, values)
}

func TestPermutationsEdges(t *testing.T) {
	assert.Exactly(t, [][]string{{}}, Slice(Permutations([]string{"A", "B"}, 0)))
	assert.Exactly(t, [][]string{}, Slice(Permutations([]string{"A", "B"}, 3)))
	assert.Exactly(t, [][]string{}, Slice(Permutations([]string{"A"}, -1)))
}

func TestPermutationCount(t *testing.T) {
	assert.Exactly(t, int64(42), PermutationCount(7, 2))
	assert.Exactly(t, int64(59280), PermutationCount(40, 3))
	assert.Exactly(t, int64(78960960), PermutationCount(40, 5))
}
