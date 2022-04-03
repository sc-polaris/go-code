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

func main() {
	defer ot.Flush()
	for {
		var n int
		fmt.Fscan(in, &n)
		if n == 0 {
			break
		}
		h := make([]int, n)
		l := make([]int, n)
		r := make([]int, n)
		for i := range h {
			fmt.Fscan(in, &h[i])
		}

		q := make([]int, 0)
		for i := 0; i < n; i++ {
			for len(q) > 0 && h[q[len(q)-1]] >= h[i] {
				q = q[:len(q)-1]
			}
			if len(q) == 0 {
				l[i] = -1
			} else {
				l[i] = q[len(q)-1]
			}
			q = append(q, i)
		}

		q = make([]int, 0)
		for i := n - 1; i >= 0; i-- {
			for len(q) > 0 && h[q[len(q)-1]] >= h[i] {
				q = q[:len(q)-1]
			}
			if len(q) == 0 {
				r[i] = n
			} else {
				r[i] = q[len(q)-1]
			}
			q = append(q, i)
		}

		var res int
		for i := 0; i < n; i++ {
			res = max(res, h[i]*(r[i]-l[i]-1))
		}
		fmt.Fprintln(ot, res)
	}
}
