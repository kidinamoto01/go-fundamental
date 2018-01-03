package main

import (
	"fmt"
)

func main() {

	a := [2]int{1, 2}
	b := [2]int{1, 2}

	//严格比较，只有大小相同才能比较类型
	fmt.Println(a == b)

	//冒泡排序

	m := [...]int{5, 8, 1, 4, 9, 2}

	fmt.Println(m)

	for i := 0; i < len(m); i++ {

		for j := i + 1; j < len(m); j++ {

			if m[i] > m[j] {
				tmp := m[i]
				m[i] = m[j]
				m[j] = tmp
			}
		}
	}

	fmt.Println(m)
}
