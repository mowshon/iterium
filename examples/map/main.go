package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	squares := iterium.Map(iterium.Range(1, 6), func(value int) int {
		return value * value
	})
	fmt.Println(iterium.Slice(squares))
}
