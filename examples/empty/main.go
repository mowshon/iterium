package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	values := iterium.Slice(iterium.Empty[int]())
	fmt.Println(values)
}
