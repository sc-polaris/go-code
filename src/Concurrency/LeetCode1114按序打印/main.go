package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)

type Foo struct {
}

func (f *Foo) first() {
	fmt.Println("first")
	ch1 <- 1
}

func (f *Foo) second() {
	<-ch1
	fmt.Println("second")
	ch2 <- 1
}

func (f *Foo) third() {
	<-ch2
	fmt.Println("third")
}

func main() {
	f := &Foo{}
	go f.first()
	go f.second()
	go f.third()
	time.Sleep(time.Second) // 防止主线程退出，子线程没运行完
}
