package main

import (
	"fmt"
)

func main() {
	input := make(map[int]string)

	input = map[int]string{1: "a", 2: "b"}

	fmt.Println(input)

	output := make(map[string]int)

	for v, k := range input {

		output[k] = v
	}

	fmt.Println(output)
}
