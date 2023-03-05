package iterium

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCombinations(t *testing.T) {
	combinations := Combinations([]string{"A", "B", "C", "D"}, 2)
	merged := Map(combinations, func(val []string) string {
		return strings.Join(val, "")
	})

	if slice, err := merged.Slice(); assert.Nil(t, err) {
		expected := []string{"AB", "AC", "AD", "BC", "BD", "CD"}
		assert.Exactly(t, expected, slice)
		assert.Equal(t, int64(6), combinations.Count())
	}
}

func TestCombinationsCount(t *testing.T) {
	assert.Exactly(t, int64(6), CombinationsCount(4, 2))
	assert.Exactly(t, int64(171), CombinationsCount(19, 2))
	assert.Exactly(t, int64(658008), CombinationsCount(40, 5))
}

func TestCombinationsEmpty(t *testing.T) {
	permutations := Combinations([]string{"A", "B"}, 10)
	assert.Exactly(t, int64(0), permutations.Count())
}
