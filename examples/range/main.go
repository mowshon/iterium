package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	fmt.Println(iterium.Slice(iterium.Range(0, 10, 2)))
	fmt.Println(iterium.Slice(iterium.Range(-3)))
}
