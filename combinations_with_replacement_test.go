package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationsWithReplacement(t *testing.T) {
	values := joinStrings(CombinationsWithReplacement([]string{"A", "B", "C", "D"}, 2))

	assert.Exactly(t, []string{
		"AA", "AB", "AC", "AD", "BB", "BC", "BD", "CC", "CD", "DD",
	}, values)
}

func TestCombinationsWithReplacementEdges(t *testing.T) {
	assert.Exactly(t, [][]string{{}}, Slice(CombinationsWithReplacement([]string{}, 0)))
	assert.Exactly(t, [][]string{}, Slice(CombinationsWithReplacement([]string{}, 1)))
	assert.Exactly(t, [][]string{}, Slice(CombinationsWithReplacement([]string{"A"}, -1)))
}

func TestCombinationsWithReplacementCount(t *testing.T) {
	assert.Exactly(t, int64(10), CombinationsWithReplacementCount(4, 2))
	assert.Exactly(t, int64(120), CombinationsWithReplacementCount(15, 2))
	assert.Exactly(t, int64(3365856), CombinationsWithReplacementCount(26, 7))
	assert.Exactly(t, int64(1), CombinationsWithReplacementCount(0, 0))
}
