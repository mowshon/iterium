package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	for number, letter := range iterium.Product2([]int{1, 2}, []string{"a", "b"}) {
		fmt.Println(number, letter)
	}
}
