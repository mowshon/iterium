package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	values := iterium.TakeWhile(iterium.Count[int](), func(value int) bool {
		return value < 5
	})
	fmt.Println(iterium.Slice(values))
}
