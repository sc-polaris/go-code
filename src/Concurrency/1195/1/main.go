package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type FizzBuzz struct {
	n          int
	numCh      chan struct{}
	fizzCh     chan struct{}
	buzzCh     chan struct{}
	fizzBuzzCh chan struct{}
	q          chan struct{}
}

// n%3==0
func (f *FizzBuzz) fizz() {
	defer wg.Done()
	for {
		select {
		case <-f.fizzCh:
			fmt.Println("fizz")
			f.numCh <- struct{}{}
		case <-f.q:
			return
		}
	}
}

// n%5==0
func (f *FizzBuzz) buzz() {
	defer wg.Done()
	for {
		select {
		case <-f.buzzCh:
			fmt.Println("buzz")
			f.numCh <- struct{}{}
		case <-f.q:
			return
		}
	}
}

// n%3 == 0 && n%5 == 0
func (f *FizzBuzz) fizzbuzz() {
	defer wg.Done()
	for {
		select {
		case <-f.fizzBuzzCh:
			fmt.Println("fizzbuzz")
			f.numCh <- struct{}{}
		case <-f.q:
			return
		}
	}
}

func (f *FizzBuzz) number() {
	defer wg.Done()
	for i := 1; i <= f.n; i++ {
		<-f.numCh

		if i%3 == 0 && i%5 == 0 {
			f.fizzBuzzCh <- struct{}{}
			continue
		}
		if i%3 == 0 {
			f.fizzCh <- struct{}{}
			continue
		}
		if i%5 == 0 {
			f.buzzCh <- struct{}{}
			continue
		}
		fmt.Println(i)
		f.numCh <- struct{}{}
	}
	f.q <- struct{}{}
	f.q <- struct{}{}
	f.q <- struct{}{}
}

func main() {
	f := &FizzBuzz{
		n:          15,
		numCh:      make(chan struct{}, 1),
		fizzCh:     make(chan struct{}),
		buzzCh:     make(chan struct{}),
		fizzBuzzCh: make(chan struct{}),
		q:          make(chan struct{}),
	}

	wg.Add(4)
	go f.fizz()
	go f.buzz()
	go f.fizzbuzz()
	go f.number()
	f.numCh <- struct{}{}

	wg.Wait()
}
