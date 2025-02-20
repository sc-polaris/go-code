package main

import (
	"fmt"
	"strconv"
	"strings"
)

func diningPhilosophers(n int) [][]string {
	ch := make(chan string)
	var res [][]string
	for i := 0; i < 5; i++ {
		go philosopher(i, n, ch)
	}
	for i := 0; i < 25*n; i++ {
		action := <-ch
		actions := strings.Split(action, ":")
		res = append(res, []string{actions[0], actions[1], actions[2]})
	}
	return res
}

func philosopher(index int, n int, ch chan string) {
	for i := 0; i < n; i++ {
		leftFork, rightFork := "1", "2"
		pick, put, eat := "1", "2", "3"
		indexStr := strconv.Itoa(index)

		ch <- indexStr + ":" + leftFork + ":" + pick
		ch <- indexStr + ":" + rightFork + ":" + pick
		ch <- indexStr + ":" + "0" + ":" + eat
		ch <- indexStr + ":" + leftFork + ":" + put
		ch <- indexStr + ":" + rightFork + ":" + put
	}
}

func main() {
	n := 1
	fmt.Println(diningPhilosophers(n))
}
