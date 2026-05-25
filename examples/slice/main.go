package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	values := iterium.Slice(iterium.Range(5))
	fmt.Println(values)
}
