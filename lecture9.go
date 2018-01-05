package main

import (
	"fmt"
)

func main() {
	a := 1
	A(&a)
	fmt.Println(a)
	b := B
	b()

	//匿名函数
	c := func() {
		fmt.Println("anonymous")
	}
	c()

	//闭包函数
	f := closure(10)
	fmt.Println(f(1)) //x=10 y=1
}

func A(i *int) {

	*i = 2

	fmt.Println(*i)

}

func B() {
	fmt.Println("func B")
}

func closure(x int) func(int) int {

	return func(y int) int {
		return x + y
	}
}
