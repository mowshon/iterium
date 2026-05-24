package iterium

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkDecryptMD5Hash(b *testing.B) {
	passHash := "d8578edf8458ce06fbc5bb76a58c5ca4" // md5("qwerty")

	for n := 0; n < b.N; n++ {
		for passLength := range Range(1, 7) {
			fmt.Println("Password Length:", passLength, "total combinations:", ProductCount(len(AsciiLowercase), passLength))

			join := func(product []string) string {
				return strings.Join(product, "")
			}

			sameHash := func(rawPassword string) bool {
				hash := md5.Sum([]byte(rawPassword))
				return hex.EncodeToString(hash[:]) == passHash
			}

			if result, ok := FirstTrue(Map(Product(AsciiLowercase, passLength), join), sameHash); ok {
				fmt.Println("Raw password:", result)
				break
			}
		}
	}
}
