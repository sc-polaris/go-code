package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	sum := 0
	for i := 1; i <= 1000000000; i++ {
		sum += 1
	}

	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
	fmt.Println(end.Sub(start).Milliseconds())
	fmt.Println(end.Sub(start).Microseconds())
}
