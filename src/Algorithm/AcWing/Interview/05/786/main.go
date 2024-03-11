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

func quickSelect(q []int, l, r, k int) int {
	if l >= r {
		return q[k]
	}

	i, j, x := l-1, r+1, q[(l+r)>>1]
	for i < j {
		for i = i + 1; q[i] < x; i++ {
		}
		for j = j - 1; q[j] > x; j-- {
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	if j >= k {
		return quickSelect(q, l, j, k)
	} else {
		return quickSelect(q, j+1, r, k)
	}
}

func main() {
	defer ot.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	q := make([]int, n)

	for i := range q {
		fmt.Fscan(in, &q[i])
	}

	fmt.Fprintln(ot, quickSelect(q, 0, n-1, k-1))
}
