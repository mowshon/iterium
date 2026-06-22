package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	values := iterium.Slice(iterium.Range(5))
	fmt.Println(values)
}
