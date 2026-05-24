package iterium

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"
)

func drainIter[T any](b *testing.B, iter Iter[T]) int {
	b.Helper()

	count := 0
	for {
		_, err := iter.Next()
		if err != nil {
			return count
		}
		count++
	}
}

func BenchmarkRangeMapFilterLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter := Filter(
			Map(Range(0, 1_000_000), func(v int) int {
				return v + 1
			}),
			func(v int) bool {
				return v%2 == 0
			},
		)

		if count := drainIter(b, iter); count != 500_000 {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkProduct26Repeat4Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter := Product(AsciiLowercase, 4)
		if count := drainIter(b, iter); count != int(iter.Count()) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkProductMapFirstTrueMD5Large(b *testing.B) {
	passHash := "02c425157ecd32f259548b33402ff6d3" // md5("zzzz")

	for n := 0; n < b.N; n++ {
		product := Product(AsciiLowercase, 4)
		joined := Map(product, func(value []string) string {
			return strings.Join(value, "")
		})
		found := FirstTrue(joined, func(rawPassword string) bool {
			hash := md5.Sum([]byte(rawPassword))
			return hex.EncodeToString(hash[:]) == passHash
		})

		result, err := found.Next()
		if err != nil {
			b.Fatal(err)
		}
		if result != "zzzz" {
			b.Fatalf("unexpected result: %s", result)
		}
	}
}

func BenchmarkCombinations26Choose5Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter := Combinations(AsciiLowercase, 5)
		if count := drainIter(b, iter); count != int(iter.Count()) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkCombinationsWithReplacement26Choose5Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter := CombinationsWithReplacement(AsciiLowercase, 5)
		if count := drainIter(b, iter); count != int(iter.Count()) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkPermutations10Pick5Large(b *testing.B) {
	symbols := AsciiLowercase[:10]

	for n := 0; n < b.N; n++ {
		iter := Permutations(symbols, 5)
		if count := drainIter(b, iter); count != int(iter.Count()) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}
