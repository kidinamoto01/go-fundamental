package main

import (
	"fmt"
)

type TZ int

func (t *TZ) Increase(num int) {
	*t += TZ(num)
}
func main() {
	var test TZ

	test.Increase(100)
	fmt.Println(test)
}
