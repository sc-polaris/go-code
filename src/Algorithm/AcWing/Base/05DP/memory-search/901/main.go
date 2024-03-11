package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 310

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	h, f [N][N]int
	dx   = [4]int{0, 1, 0, -1}
	dy   = [4]int{1, 0, -1, 0}
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dp(x, y int) int {
	if f[x][y] != 0 { // 如果这里被遍历过一次，就直接返回
		return f[x][y]
	}

	f[x][y] = 1
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if 0 <= a && a < n && 0 <= b && b < m && h[a][b] < h[x][y] {
			f[x][y] = max(f[x][y], dp(a, b)+1)
		}
	}

	return f[x][y]
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &h[i][j])
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res = max(res, dp(i, j))
		}
	}

	fmt.Fprintln(ot, res)
}
