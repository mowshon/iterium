package iternium

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestProduct(t *testing.T) {
	product := Product([]string{"A", "B", "C", "D"}, 2)
	merged := Map(product, func(val []string) string {
		return strings.Join(val, "")
	})

	if slice, err := merged.Slice(); assert.Nil(t, err) {
		expected := []string{
			"AA", "AB", "AC", "AD", "BA",
			"BB", "BC", "BD", "CA", "CB",
			"CC", "CD", "DA", "DB", "DC", "DD",
		}

		assert.Exactly(t, expected, slice)
		assert.Exactly(t, int64(16), product.Count())
	}
}

func TestProductInteger(t *testing.T) {
	product := Product([]int{0, 1}, 2)
	if slice, err := product.Slice(); assert.Nil(t, err) {
		expected := [][]int{
			{0, 0}, {0, 1}, {1, 0}, {1, 1},
		}

		assert.Exactly(t, expected, slice)
	}
}

func TestProductCount(t *testing.T) {
	assert.Exactly(t, int64(16), ProductCount(4, 2))
	assert.Exactly(t, int64(1600), ProductCount(40, 2))
	assert.Exactly(t, int64(20511149), ProductCount(29, 5))
}
