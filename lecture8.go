package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	m = map[int]string{}

	fmt.Println(m)

	n := make(map[int]string)
	fmt.Println(n)

	m[0] = "a"
	fmt.Println(m)
	delete(m, 0)
	fmt.Println(m)

	var o map[int]map[int]string
	o = make(map[int]map[int]string)
	o[1] = make(map[int]string)
	o[1][1] = "OK"
	fmt.Println(o)

	a, ok := o[2][1]
	fmt.Println(a, ok)
	if !ok {
		o[2] = make(map[int]string)
	}

	o[2][1] = "GOOD"
	a, ok = o[2][1]
	fmt.Println(a, ok)
}
