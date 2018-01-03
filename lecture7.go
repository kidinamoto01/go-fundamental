package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//make([]T,len,cap)
	s1 := a[2:4]

	fmt.Println(s1, len(s1), cap(s1))

	s2 := a[1:5]

	fmt.Println(s2, len(s2), cap(s2))

	s2[1] = 9
	fmt.Println(s1)

	s2 = append(s2, 0)
	fmt.Println(s2, len(s2), cap(s2))

	copy(s1, s2)
	fmt.Println(s1)
}
