package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	values := joinStrings(Product([]string{"A", "B", "C", "D"}, 2))

	assert.Exactly(t, []string{
		"AA", "AB", "AC", "AD", "BA",
		"BB", "BC", "BD", "CA", "CB",
		"CC", "CD", "DA", "DB", "DC", "DD",
	}, values)
}

func TestProductInteger(t *testing.T) {
	assert.Exactly(t, [][]int{
		{0, 0}, {0, 1}, {1, 0}, {1, 1},
	}, Slice(Product([]int{0, 1}, 2)))
}

func TestProductEdges(t *testing.T) {
	assert.Exactly(t, [][]string{{}}, Slice(Product([]string{}, 0)))
	assert.Exactly(t, [][]string{}, Slice(Product([]string{}, 1)))
	assert.Exactly(t, [][]string{}, Slice(Product([]string{"A"}, -1)))
}

func TestProduct2(t *testing.T) {
	values := Slice2(Product2([]int{1, 2}, []string{"a", "b"}))

	assert.Exactly(t, []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 1, Second: "b"},
		{First: 2, Second: "a"},
		{First: 2, Second: "b"},
	}, values)
}

func TestChan2(t *testing.T) {
	ctx, cancel := testContext()
	defer cancel()

	var values []Pair[int, string]
	for value := range Chan2(ctx, Product2([]int{1}, []string{"a", "b"})) {
		values = append(values, value)
	}

	assert.Exactly(t, []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 1, Second: "b"},
	}, values)
}

func TestProductCount(t *testing.T) {
	assert.Exactly(t, int64(16), ProductCount(4, 2))
	assert.Exactly(t, int64(1600), ProductCount(40, 2))
	assert.Exactly(t, int64(20511149), ProductCount(29, 5))
}

func TestProductIntoReusesBuffer(t *testing.T) {
	var first []byte

	ProductBytesInto([]byte("ab"), 2, func(value []byte) bool {
		if first == nil {
			first = value
		}
		return false
	})

	assert.Exactly(t, []byte("aa"), first)
}

func TestProductStringInto(t *testing.T) {
	var values []string

	ProductStringInto("ab", 2, func(value []byte) bool {
		values = append(values, string(value))
		return true
	})

	assert.Exactly(t, []string{"aa", "ab", "ba", "bb"}, values)
}
