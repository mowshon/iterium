package main

import (
	"crypto/md5"
	"fmt"
	"time"

	iterium "github.com/mowshon/iterium"
)

func main() {
	password := "qwerty"
	targetHash := md5.Sum([]byte(password))
	alphabet := iterium.AsciiLowercaseBytes
	maxLength := len(password)

	var total int64
	for length := 1; length <= maxLength; length++ {
		count, ok := iterium.ProductCountOK(len(alphabet), length)
		if !ok {
			panic("combination count overflowed")
		}
		total += count
	}

	var checked int64
	var found string

	started := time.Now()
	for length := 1; length <= maxLength; length++ {
		iterium.ProductBytesInto(alphabet, length, func(candidate []byte) bool {
			checked++
			hash := md5.Sum(candidate)
			if hash == targetHash {
				found = string(candidate)
				return false
			}
			return true
		})

		if found != "" {
			break
		}
	}
	elapsed := time.Since(started)

	fmt.Printf("target md5: %x\n", targetHash)
	fmt.Printf("alphabet: %q (%d symbols)\n", string(alphabet), len(alphabet))
	fmt.Printf("max length: %d\n", maxLength)
	fmt.Printf("total combinations: %d\n", total)
	fmt.Printf("checked combinations: %d\n", checked)
	fmt.Printf("found password: %q\n", found)
	fmt.Printf("elapsed seconds: %.6f\n", elapsed.Seconds())
}
