package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	odds := iterium.FilterFalse(iterium.Range(10), func(value int) bool {
		return value%2 == 0
	})
	fmt.Println(iterium.Slice(odds))
}
