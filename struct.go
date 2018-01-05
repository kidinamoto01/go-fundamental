package main

import (
	"fmt"
)

type person struct {
	Name string //必须大写
	Age  int
}

type teacher struct {
	person
	rank string
}
type employee struct {
	Name     string
	Age      int
	Contract struct {
		Phone, city string
	}
}

func main() {
	//推荐在给结构赋值的时候采用指针的方式
	a := &person{
		Name: "joe",
		Age:  19, //不要忘记逗号
	}
	fmt.Println(a)

	a.Name = "OK"
	fmt.Println(a)
	A(a)
	fmt.Println(a)

	//匿名结构
	test := struct {
		Age  int
		Name string
	}{
		Age:  10,
		Name: "Man",
	}

	fmt.Println(test)

	//匿名结构的赋值
	e := employee{Name: "Cat", Age: 20}
	e.Contract.Phone = "12345678"
	e.Contract.city = "NY"

	fmt.Println(e)

	//结构的嵌套
	t := teacher{rank: "2"}
	t.Age = 30
	t.Name = "tom"

	fmt.Println(t)

}

func A(per *person) {
	per.Age = 13
}
