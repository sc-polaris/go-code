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

/*
算法框架：用单调栈找每个障碍物左边第一个比它高的位置
*/

func main() {
	defer ot.Flush()

	var n int
	fmt.Fscan(in, &n)

	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}

	q := make([]int, 0)
	res := 0
	for i := range h {
		last := 0
		for len(q) > 0 && h[q[len(q)-1]] <= h[i] {
			res += (h[q[len(q)-1]] - last) * (i - q[len(q)-1] - 1)
			last = h[q[len(q)-1]]
			q = q[:len(q)-1]
		}
		if len(q) > 0 {
			res += (h[i] - last) * (i - q[len(q)-1] - 1)
		}
		q = append(q, i)
	}

	fmt.Fprintln(ot, res)
}
