package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
	n   = 5
)

func zero() {
	for i := 1; i <= n; i++ {
		<-ch3
		fmt.Print(0)
		if i&1 == 1 {
			ch1 <- 1
		} else {
			ch2 <- 1
		}
	}
}

func odd() {
	for i := 1; i <= n; i += 2 {
		<-ch1
		fmt.Print(i)
		ch3 <- 1
	}
}

func even() {
	for i := 2; i <= n; i += 2 {
		<-ch2
		fmt.Print(i)
		ch3 <- 1
	}
}

func main() {
	go zero()
	go odd()
	go even()
	ch3 <- 1                // 3个线程启动后才开始输出
	time.Sleep(time.Second) // 防止主线程退出后，子线程没运行完毕
}
