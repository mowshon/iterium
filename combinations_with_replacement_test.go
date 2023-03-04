package iternium

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCombinationsWithReplacement(t *testing.T) {
	withReplacement := CombinationsWithReplacement([]string{"A", "B", "C", "D"}, 2)
	merged := Map(withReplacement, func(val []string) string {
		return strings.Join(val, "")
	})

	if slice, err := merged.Slice(); assert.Nil(t, err) {
		expected := []string{
			"AA", "AB", "AC", "AD", "BB", "BC", "BD", "CC", "CD", "DD",
		}

		assert.Exactly(t, expected, slice)
		assert.Exactly(t, int64(10), withReplacement.Count())
	}
}

func TestCombinationsWithReplacementCount(t *testing.T) {
	assert.Exactly(t, int64(10), CombinationsWithReplacementCount(4, 2))
	assert.Exactly(t, int64(120), CombinationsWithReplacementCount(15, 2))
	assert.Exactly(t, int64(3365856), CombinationsWithReplacementCount(26, 7))
}
