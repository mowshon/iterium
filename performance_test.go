package iterium

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"
)

func BenchmarkSeqRangeMapFilterLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count := 0
		values := Filter(
			Map(Range(0, 1_000_000), func(v int) int {
				return v + 1
			}),
			func(v int) bool {
				return v%2 == 0
			},
		)

		for range values {
			count++
		}
		if count != 500_000 {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkSeqProduct26Repeat4Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count := 0
		for range Product(AsciiLowercase, 4) {
			count++
		}
		if count != int(ProductCount(len(AsciiLowercase), 4)) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkSeqProductInto26Repeat4Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count := 0
		ProductInto(AsciiLowercase, 4, func([]string) bool {
			count++
			return true
		})
		if count != int(ProductCount(len(AsciiLowercase), 4)) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkSeqProductMapFirstTrueMD5Large(b *testing.B) {
	passHash := "02c425157ecd32f259548b33402ff6d3" // md5("zzzz")

	for n := 0; n < b.N; n++ {
		joined := Map(Product(AsciiLowercase, 4), func(value []string) string {
			return strings.Join(value, "")
		})
		result, ok := FirstTrue(joined, func(rawPassword string) bool {
			hash := md5.Sum([]byte(rawPassword))
			return hex.EncodeToString(hash[:]) == passHash
		})
		if !ok {
			b.Fatal("not found")
		}
		if result != "zzzz" {
			b.Fatalf("unexpected result: %s", result)
		}
	}
}

func BenchmarkSeqProductBytesIntoMD5Large(b *testing.B) {
	passHash := md5.Sum([]byte("zzzz"))

	for n := 0; n < b.N; n++ {
		var result []byte
		ProductBytesInto(AsciiLowercaseBytes, 4, func(value []byte) bool {
			hash := md5.Sum(value)
			if bytes.Equal(hash[:], passHash[:]) {
				result = append(result[:0], value...)
				return false
			}
			return true
		})
		if string(result) != "zzzz" {
			b.Fatalf("unexpected result: %s", result)
		}
	}
}

func BenchmarkSeqCombinations26Choose5Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count := 0
		for range Combinations(AsciiLowercase, 5) {
			count++
		}
		if count != int(CombinationsCount(len(AsciiLowercase), 5)) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkSeqCombinationsWithReplacement26Choose5Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count := 0
		for range CombinationsWithReplacement(AsciiLowercase, 5) {
			count++
		}
		if count != int(CombinationsWithReplacementCount(len(AsciiLowercase), 5)) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkSeqPermutations10Pick5Large(b *testing.B) {
	symbols := AsciiLowercase[:10]

	for n := 0; n < b.N; n++ {
		count := 0
		for range Permutations(symbols, 5) {
			count++
		}
		if count != int(PermutationCount(len(symbols), 5)) {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}
