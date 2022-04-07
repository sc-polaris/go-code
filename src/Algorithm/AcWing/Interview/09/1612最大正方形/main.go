package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	f    [N][N]int
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var w int
			fmt.Fscan(in, &w)
			if w == 1 {
				f[i][j] = min(f[i-1][j-1], min(f[i-1][j], f[i][j-1])) + 1
			}
			res = max(res, f[i][j])
		}
	}

	fmt.Fprintln(ot, res*res)
}
