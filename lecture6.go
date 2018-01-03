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
	a := -1
	if a := 1; a > 0 {

		//a的作用范围只在if语句内
		fmt.Println("a>0")
	}
}
