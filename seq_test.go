package iterium

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeMapFilterSeq(t *testing.T) {
	values := FilterSeq(
		MapSeq(RangeSeq(0, 10), func(value int) int {
			return value + 1
		}),
		func(value int) bool {
			return value%2 == 0
		},
	)

	assert.Exactly(t, []int{2, 4, 6, 8, 10}, SliceSeq(values))
}

func TestEmptySeq(t *testing.T) {
	assert.Exactly(t, []int{}, SliceSeq(EmptySeq[int]()))
}

func TestAccumulateSeq(t *testing.T) {
	values := AccumulateSeq(NewSeq(1, 2, 3, 4, 5), add)
	assert.Exactly(t, []int{1, 3, 6, 10, 15}, SliceSeq(values))
}

func TestTakeWhileSeq(t *testing.T) {
	values := TakeWhileSeq(NewSeq(1, 4, 6, 4, 1), func(value int) bool {
		return value < 5
	})

	assert.Exactly(t, []int{1, 4}, SliceSeq(values))
}

func TestDropWhileSeq(t *testing.T) {
	values := DropWhileSeq(NewSeq(1, 4, 6, 4, 1), func(value int) bool {
		return value < 5
	})

	assert.Exactly(t, []int{6, 4, 1}, SliceSeq(values))
}

func TestDropWhileSeqEmptyWhenPredicateNeverFalse(t *testing.T) {
	values := DropWhileSeq(NewSeq(1, 2, 3), func(value int) bool {
		return value < 5
	})

	assert.Exactly(t, []int{}, SliceSeq(values))
}

func TestStarMapSeq(t *testing.T) {
	values := StarMapSeq(ProductSeq([]int{1, 2}, 2), func(first, second int) int {
		return first + second
	})

	assert.Exactly(t, []int{2, 3, 3, 4}, SliceSeq(values))
}

func TestCycleSeq(t *testing.T) {
	var values []int

	for value := range CycleSeq(NewSeq(1, 2, 3)) {
		values = append(values, value)
		if len(values) == 8 {
			break
		}
	}

	assert.Exactly(t, []int{1, 2, 3, 1, 2, 3, 1, 2}, values)
}

func TestCycleSeqEmpty(t *testing.T) {
	assert.Exactly(t, []int{}, SliceSeq(CycleSeq(EmptySeq[int]())))
}

func TestFirstTrueFalseSeq(t *testing.T) {
	firstTrue, ok := FirstTrueSeq(NewSeq(1, 2, 3, 4), func(value int) bool {
		return value > 2
	})
	assert.True(t, ok)
	assert.Exactly(t, 3, firstTrue)

	firstFalse, ok := FirstFalseSeq(NewSeq(1, 2, 3, 4), func(value int) bool {
		return value < 3
	})
	assert.True(t, ok)
	assert.Exactly(t, 3, firstFalse)
}

func TestProductSeq(t *testing.T) {
	values := MapSeq(ProductSeq([]string{"A", "B", "C", "D"}, 2), func(value []string) string {
		return strings.Join(value, "")
	})

	expected := []string{
		"AA", "AB", "AC", "AD", "BA",
		"BB", "BC", "BD", "CA", "CB",
		"CC", "CD", "DA", "DB", "DC", "DD",
	}
	assert.Exactly(t, expected, SliceSeq(values))
}

func TestProduct2Seq(t *testing.T) {
	values := Slice2(Product2Seq([]int{1, 2}, []string{"a", "b"}))

	assert.Exactly(t, []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 1, Second: "b"},
		{First: 2, Second: "a"},
		{First: 2, Second: "b"},
	}, values)
}

func TestChan2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var values []Pair[int, string]
	for value := range Chan2(ctx, Product2Seq([]int{1}, []string{"a", "b"})) {
		values = append(values, value)
	}

	assert.Exactly(t, []Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 1, Second: "b"},
	}, values)
}

func TestCombinationsSeq(t *testing.T) {
	values := MapSeq(CombinationsSeq([]string{"A", "B", "C", "D"}, 2), func(value []string) string {
		return strings.Join(value, "")
	})

	assert.Exactly(t, []string{"AB", "AC", "AD", "BC", "BD", "CD"}, SliceSeq(values))
}

func TestCombinationsWithReplacementSeq(t *testing.T) {
	values := MapSeq(CombinationsWithReplacementSeq([]string{"A", "B", "C", "D"}, 2), func(value []string) string {
		return strings.Join(value, "")
	})

	assert.Exactly(t, []string{
		"AA", "AB", "AC", "AD", "BB", "BC", "BD", "CC", "CD", "DD",
	}, SliceSeq(values))
}

func TestPermutationsSeq(t *testing.T) {
	values := MapSeq(PermutationsSeq([]string{"A", "B", "C", "D"}, 2), func(value []string) string {
		return strings.Join(value, "")
	})

	assert.Exactly(t, []string{
		"AB", "AC", "AD", "BA", "BC", "BD",
		"CA", "CB", "CD", "DA", "DB", "DC",
	}, SliceSeq(values))
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
