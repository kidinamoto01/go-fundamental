/*
在go语言中不支持指针，通过&来访问地址，＊来访问值
默认值为nil，而不是NULL
*/

package main

import (
	"fmt"
)

func main() {

	//指针类型在Go语言中的应用
	a := 1
	var p *int = &a

	fmt.Println(*p)

	//递增递减，只能作为语句，不能作为符号，只能放在右边

	//判断语句if
	//条件表达式没有括号
	m := -1
	if m := 1; m > 0 {

		//a的作用范围只在if语句内
		fmt.Println("m>0")
	}
	fmt.Println(m)
	//for实现多个关键字的形式

	b := 1
	//无限循环
	for {

		b++
		if b > 3 {
			break
		}

		fmt.Println(b)
	}

	fmt.Println("over")

	//与java相类似的for循环形式
	s := "string"
	for i := 0; i < len(s); i++ {

		fmt.Println(i)
	}

	//单一判断语句
	i := 0
	for i < len(s) {
		i++
		fmt.Println(i)
	}

	//switch 语句的多种形式
	switch i {
	case 0:
		fmt.Println("i=0")
	case 1:
		fmt.Println("i=1")
	default:
		fmt.Println("i>1")

	}

	//switch后没有条件表达式
	switch {
	case i > 1:
		fmt.Println("i>1")
		fallthrough //不执行break，继续判断
	case i > 2:
		fmt.Println("i>2")

	}

}
