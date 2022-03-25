package main

import (
	"fmt"
	"time"
)

type Foo struct {
	firstChan  chan int
	secondChan chan int
}

func (f *Foo) first() {
	fmt.Println("first")
	f.firstChan <- 1
}

func (f *Foo) second() {
	<-f.firstChan
	fmt.Println("second")
	f.secondChan <- 1
}

func (f *Foo) third() {
	<-f.secondChan
	fmt.Println("third")
}

func main() {
	f := &Foo{
		firstChan:  make(chan int),
		secondChan: make(chan int),
	}
	go f.first()
	go f.second()
	go f.third()
	time.Sleep(time.Second) // 防止主线程退出，子线程没运行完
}
