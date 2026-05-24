package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	value, ok := iterium.FirstFalse(iterium.Range(10), func(value int) bool {
		return value < 5
	})
	fmt.Println(value, ok)
}
