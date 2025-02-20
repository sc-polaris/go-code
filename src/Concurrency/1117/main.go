package main

import (
	"fmt"
	"time"
)

var hydrogen_chan1 = make(chan int)
var hydrogen_chan2 = make(chan int)
var oxygen_chan = make(chan int)

func hydrogen1(n int) {
	for i := 0; i < n; i++ {
		<-hydrogen_chan1
		fmt.Printf("H")
		hydrogen_chan2 <- 1
	}
}
func hydrogen2(n int) {
	for i := 0; i < n; i++ {
		<-hydrogen_chan2
		fmt.Printf("H")
		oxygen_chan <- 1
	}
}
func oxygen(n int) {
	for i := 0; i < n; i++ {
		<-oxygen_chan
		fmt.Println("O")
		hydrogen_chan1 <- 1
	}
}

func main() {
	water := "HOHHHOOHHOOHHHH"
	n := len(water) / 3
	go hydrogen1(n)
	go hydrogen2(n)
	go oxygen(n)
	hydrogen_chan1 <- 1
	time.Sleep(time.Second)
}
