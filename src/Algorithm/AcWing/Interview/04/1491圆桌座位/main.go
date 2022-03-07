package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 11

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	g    [N][N]bool // g表示朋友关系
	st   [N]bool    // st表示人是否使用过的状态
	pos  [N]int     // 每个位置上是谁
)

func dfs(u int) int {
	if u > n { // n个人已经排满
		if g[pos[n]][pos[1]] { // 判断最后一人是否与第一人为朋友关系
			return 0
		}
		return 1
	}

	var res int
	for i := 2; i <= n; i++ { // 题干中给出人是从编号1开始的,第一个人已经固定从2开始搜索
		if !st[i] && !g[i][pos[u-1]] { // 该人没用过且与前一个人非朋友关系
			pos[u] = i
			st[i] = true
			res += dfs(u + 1)
			st[i] = false // 还原现场
		}
	}

	return res
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	for ; m > 0; m-- {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a][b] = true // 朋友关系初始化
		g[b][a] = true
	}

	pos[1] = 1 // 第一个人挂上去
	st[1] = true

	fmt.Fprintln(ot, dfs(2))
}
