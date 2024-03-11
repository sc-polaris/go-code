package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 15

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
	x  []int // x[i] 表示第i列皇后的横坐标, i表示纵坐标
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(u int) bool {
	for i := 0; i < u; i++ {
		// 前者判断是否处于同一对角线。后者判断是否属于一行
		if abs(i-u) == abs(x[i]-x[u]) || x[i] == x[u] {
			return false
		}
	}

	return true
}

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if j == x[i] {
					fmt.Fprintf(ot, "Q")
				} else {
					fmt.Fprintf(ot, ".")
				}
			}
			fmt.Fprintln(ot)
		}
		fmt.Fprintln(ot)
		return
	}

	for i := 0; i < n; i++ {
		x[u] = i
		if check(u) {
			dfs(u + 1)
		}
	}
}

func main() {
	defer ot.Flush()
	fmt.Fscan(in, &n)
	x = make([]int, n)
	dfs(0)
}
