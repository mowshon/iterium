package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	values := iterium.Slice(iterium.Empty[int]())
	fmt.Println(values)
}
