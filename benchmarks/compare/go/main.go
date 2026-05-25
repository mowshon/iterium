package main

import (
	"crypto/md5"
	"fmt"
	"os"

	"github.com/mowshon/iterium"
)

const (
	rangeLimit      = 10_000_000
	productRepeat   = 5
	md5Password     = "zzzzz"
	combinationN    = 35
	combinationR    = 8
	permutationN    = 12
	permutationR    = 8
	replacementN    = 30
	replacementR    = 8
	expectedRange   = rangeLimit / 2
	expectedProduct = 11_881_376
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go-compare <range-map-filter|product-repeat5|md5-repeat5|combinations|combinations-with-replacement|permutations>")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "range-map-filter":
		count := 0
		values := iterium.Filter(
			iterium.Map(iterium.Range(0, rangeLimit), func(v int) int {
				return v + 1
			}),
			func(v int) bool {
				return v%2 == 0
			},
		)
		for range values {
			count++
		}
		mustEqual(count, expectedRange)
		fmt.Println(count)
	case "product-repeat5":
		count := 0
		iterium.ProductBytesInto(iterium.AsciiLowercaseBytes, productRepeat, func([]byte) bool {
			count++
			return true
		})
		mustEqual(count, expectedProduct)
		fmt.Println(count)
	case "md5-repeat5":
		target := md5.Sum([]byte(md5Password))
		var found []byte
		iterium.ProductBytesInto(iterium.AsciiLowercaseBytes, productRepeat, func(candidate []byte) bool {
			hash := md5.Sum(candidate)
			if hash == target {
				found = append(found[:0], candidate...)
				return false
			}
			return true
		})
		mustEqual(string(found), md5Password)
		fmt.Println(string(found))
	case "combinations":
		count := 0
		iterium.CombinationsInto(symbols(combinationN), combinationR, func([]int) bool {
			count++
			return true
		})
		mustEqual(count, int(iterium.CombinationsCount(combinationN, combinationR)))
		fmt.Println(count)
	case "combinations-with-replacement":
		count := 0
		iterium.CombinationsWithReplacementInto(symbols(replacementN), replacementR, func([]int) bool {
			count++
			return true
		})
		mustEqual(count, int(iterium.CombinationsWithReplacementCount(replacementN, replacementR)))
		fmt.Println(count)
	case "permutations":
		count := 0
		iterium.PermutationsInto(symbols(permutationN), permutationR, func([]int) bool {
			count++
			return true
		})
		mustEqual(count, int(iterium.PermutationCount(permutationN, permutationR)))
		fmt.Println(count)
	default:
		fmt.Fprintf(os.Stderr, "unknown benchmark: %s\n", os.Args[1])
		os.Exit(2)
	}
}

func symbols(n int) []int {
	values := make([]int, n)
	for i := range values {
		values[i] = i
	}
	return values
}

func mustEqual[T comparable](got, want T) {
	if got != want {
		fmt.Fprintf(os.Stderr, "unexpected result: got %v, want %v\n", got, want)
		os.Exit(1)
	}
}
