package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	values := iterium.SliceUntil(iterium.Count[int](), func(value int) bool {
		return value == 5
	})
	fmt.Println(values)
}
