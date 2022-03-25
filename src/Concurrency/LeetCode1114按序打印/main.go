package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)

func first() {
	fmt.Println("1")
	ch1 <- 1
}

func second() {
	<-ch1
	fmt.Println("2")
	ch2 <- 1
}

func third() {
	<-ch2
	fmt.Println("3")
}

func main() {
	go first()
	go second()
	go third()
	time.Sleep(time.Second) // 防止主线程退出，子线程没运行完
}
