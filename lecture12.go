package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go ")
		c <- true
		close(c) //必须明确关闭，否则会死锁
	}()

	//读出channel内容后结束main函数
	for v := range c {
		fmt.Println(v)
	}
}
