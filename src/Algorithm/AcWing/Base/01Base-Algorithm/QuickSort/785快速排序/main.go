package main

import (
	"fmt"
	"bufio"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}

	i, j, x := l-1, r+1, q[(l+r)>>1]
	for i < j {
		for i = i+1; q[i] < x; i++{
		}
		for j = j-1; q[j] > x; j--{
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func main() {
    defer ot.Flush()

	var n int
	fmt.Fscan(in, &n)

	q := make([]int, n)
	for i := range q {
		fmt.Fscan(in, &q[i])
	}

	quickSort(q, 0, n-1)

	for _, v := range q {
		fmt.Fprintf(ot, "%d ", v)
	}
}