package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	value, ok := iterium.FirstTrue(iterium.Range(10), func(value int) bool {
		return value > 5
	})
	fmt.Println(value, ok)
}
