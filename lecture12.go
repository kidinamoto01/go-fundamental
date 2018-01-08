package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go ")
		c <- true
	}()

	<-c //读出channel内容后结束main函数
}
