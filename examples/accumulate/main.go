package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	totals := iterium.Accumulate(iterium.New(1, 2, 3, 4), func(left, right int) int {
		return left + right
	})
	fmt.Println(iterium.Slice(totals))
}
