package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	var n int
	fmt.Fscan(in, &n)

	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}

	res := 0
	i, j := 0, n-1
	for i < j {
		res = max(res, min(h[i], h[j])*(j-i))
		if h[i] < h[j] {
			i++
		} else {
			j--
		}
	}

	fmt.Fprintln(ot, res)
}
