package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	for value := range iterium.New("red", "green", "blue") {
		fmt.Println(value)
	}
}
