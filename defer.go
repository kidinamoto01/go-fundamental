package main

import (
	"fmt"
)

func main() {
	//相反顺序逐个执行
	//先进后出
	fmt.Println("a")
	//得到 a,c b
	defer fmt.Println("b")
	defer fmt.Println("c")

	// //循环中的普通调用
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	//循环中的函数调用,作为闭包
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	//在发生严重错误时候，在defer 中执行recover函数

	A()
	B()
	C()
}

func A() {

	fmt.Println("func A")
}

func B() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()

	panic("panic in B")
}

func C() {

	fmt.Println("func C")
}
