package main

import (
	"fmt"
	"time"
)

var fizzChan = make(chan int)
var buzzChan = make(chan int)
var fizzbuzzChan = make(chan int)
var stopChan = make(chan int)

func fizz() {
	for {
		if _, ok := <-fizzChan; ok {
			fmt.Println("fizz")
			stopChan <- 1
		} else {
			break
		}
	}
}

func buzz() {
	for {
		if _, ok := <-buzzChan; ok {
			fmt.Println("buzz")
			stopChan <- 1
		} else {
			break
		}
	}
}

func fizzbuzz() {
	for {
		if _, ok := <-fizzbuzzChan; ok {
			fmt.Println("fizzbuzz")
			stopChan <- 1
		} else {
			break
		}
	}
}

func work(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fizzbuzzChan <- 1
			<-stopChan
		} else if i%3 == 0 {
			fizzChan <- 1
			<-stopChan
		} else if i%5 == 0 {
			buzzChan <- 1
			<-stopChan
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	go fizz()
	go buzz()
	go fizzbuzz()
	n := 35
	work(n)
	time.Sleep(1 * time.Second)
}
