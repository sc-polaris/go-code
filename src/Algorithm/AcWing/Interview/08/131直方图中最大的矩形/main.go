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
		h := make([]int, n+2)
		l := make([]int, n+1)
		r := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &h[i])
		}
		h[0], h[n+1] = -1, -1

		var q []int
		// 求左边界
		q = append(q, 0)
		for i := 1; i <= n; i++ {
			for h[q[len(q)-1]] >= h[i] {
				q = q[:len(q)-1]
			}
			l[i] = q[len(q)-1]
			q = append(q, i)
		}

		q = make([]int, 0)
		// 求右边界
		q = append(q, n+1)
		for i := n; i >= 1; i-- {
			for h[q[len(q)-1]] >= h[i] {
				q = q[:len(q)-1]
			}
			r[i] = q[len(q)-1]
			q = append(q, i)
		}

		var res int
		for i := 1; i <= n; i++ {
			res = max(res, h[i]*(r[i]-l[i]-1))
		}
		fmt.Fprintln(ot, res)
	}
}
