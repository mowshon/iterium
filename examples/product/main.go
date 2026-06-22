package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	for value := range iterium.Product([]string{"A", "B"}, 2) {
		fmt.Println(value)
	}
}
