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
	a    []int
)

func dfs(u, start int) {
	// u-1:已经选的数，n-start+1：还可以选的数，二者相加<m则剪枝
	if u+n-start < m {
		return
	}
	if u == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(ot, "%d ", a[i])
		}
		fmt.Fprintln(ot)
		return
	}

	for i := start; i < n; i++ {
		a[u] = i + 1
		dfs(u+1, i+1)
		a[u] = 0 // 还原
	}

}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	a = make([]int, m)

	dfs(0, 0)
}
