package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 15

var (
	in  = bufio.NewReader(os.Stdin)
	ot  = bufio.NewWriter(os.Stdout)
	n   int
	g   [N][N]byte
	col [N]bool     // 列
	dg  [N * 2]bool // 对角线
	udg [N * 2]bool // 反对角线
)

// 暴搜：u：行
func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Fprintf(ot, "%c", g[i][j])
			}
			fmt.Fprintln(ot)
		}
		fmt.Fprintln(ot)
		return
	}

	for i := 0; i < n; i++ {
		if !col[i] && !dg[u+i] && !udg[u-i+n] {
			col[i], dg[u+i], udg[u-i+n] = true, true, true
			g[u][i] = 'Q'
			dfs(u + 1)
			g[u][i] = '.'
			col[i], dg[u+i], udg[u-i+n] = false, false, false
		}
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = '.'
		}
	}

	dfs(0)
}
