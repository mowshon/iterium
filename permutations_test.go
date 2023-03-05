package iterium

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPermutations(t *testing.T) {
	permutations := Permutations([]string{"A", "B", "C", "D"}, 2)
	merged := Map(permutations, func(val []string) string {
		return strings.Join(val, "")
	})

	if slice, err := merged.Slice(); assert.Nil(t, err) {
		expected := []string{
			"AB", "AC", "AD", "BA", "BC", "BD",
			"CB", "CA", "CD", "DB", "DC", "DA",
		}

		assert.Exactly(t, expected, slice)
		assert.Exactly(t, int64(12), permutations.Count())
	}
}

func TestPermutationCount(t *testing.T) {
	assert.Exactly(t, int64(42), PermutationCount(7, 2))
	assert.Exactly(t, int64(59280), PermutationCount(40, 3))
	assert.Exactly(t, int64(78960960), PermutationCount(40, 5))
}

func TestPermutationsEmpty(t *testing.T) {
	permutations := Permutations([]string{"A", "B"}, 10)
	assert.Exactly(t, int64(0), permutations.Count())
}
