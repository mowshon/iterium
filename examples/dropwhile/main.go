package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	values := iterium.DropWhile(iterium.Range(10), func(value int) bool {
		return value < 5
	})
	fmt.Println(iterium.Slice(values))
}
