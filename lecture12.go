package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)

	o := make(chan bool, 2)
	//o := make(chan bool)
	//select作为接受者
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1 :", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2 :", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 2
	c2 <- "hello"

	close(c1)
	//只要关闭一个channel close(c2)
	for i := 0; i < 2; i++ {
		<-o
	}
	//select作为发送者

}
