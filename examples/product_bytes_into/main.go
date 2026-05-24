package main

import (
	"bytes"
	"crypto/md5"
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	target := md5.Sum([]byte("az"))
	var found []byte

	iterium.ProductBytesInto(iterium.AsciiLowercaseBytes, 2, func(value []byte) bool {
		hash := md5.Sum(value)
		if bytes.Equal(hash[:], target[:]) {
			found = append(found[:0], value...)
			return false
		}
		return true
	})

	fmt.Println(string(found))
}
