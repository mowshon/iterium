package iterium

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkDecryptMD5Hash(b *testing.B) {
	// result of md5("qwerty") = d8578edf8458ce06fbc5bb76a58c5ca4
	passHash := "d8578edf8458ce06fbc5bb76a58c5ca4"

	for passLength := range Range(1, 7).Chan() {
		product := Product(AsciiLowercase, passLength)
		fmt.Println("Password Length:", passLength, "total combinations:", product.Count())

		// Merge a slide into a string.
		// []string{"a", "b", "c"} => "abc"
		join := func(product []string) string {
			return strings.Join(product, "")
		}

		// Check the hash of a raw password with an unknown hash.
		sameHash := func(rawPassword string) bool {
			hash := md5.Sum([]byte(rawPassword))
			return hex.EncodeToString(hash[:]) == passHash
		}

		// Combine iterators to achieve the goal...
		decrypt := FirstTrue(
			Map(product, join), sameHash,
		)

		if result, err := decrypt.Next(); err == nil {
			fmt.Println("Raw password:", result)
			break
		}
	}
}
