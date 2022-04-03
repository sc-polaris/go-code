package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	s    [][]int
	l, r []int
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func work(h []int) int {
	q := make([]int, 0)

	h[0], h[m+1] = -1, -1
	q = append(q, 0)
	for i := 1; i <= m; i++ {
		for h[q[len(q)-1]] >= h[i] {
			q = q[:len(q)-1]
		}
		l[i] = q[len(q)-1]
		q = append(q, i)
	}

	q = make([]int, 0)
	q = append(q, m+1)
	for i := m; i >= 1; i-- {
		for h[q[len(q)-1]] >= h[i] {
			q = q[:len(q)-1]
		}
		r[i] = q[len(q)-1]
		q = append(q, i)
	}

	res := 0
	for i := 1; i <= m; i++ {
		res = max(res, h[i]*(r[i]-l[i]-1))
	}

	return res
}

func main() {
	defer ot.Flush()
	fmt.Fscan(in, &n, &m)

	s = make([][]int, n+1)
	l = make([]int, m+1)
	r = make([]int, m+1)
	for i := range s {
		s[i] = make([]int, m+2)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var c string
			fmt.Fscan(in, &c)
			if c == "F" {
				s[i][j] = s[i-1][j] + 1
			}
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		res = max(res, work(s[i]))
	}

	fmt.Fprintln(ot, res*3)
}
