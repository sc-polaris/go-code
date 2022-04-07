package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 10

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	g    = make([][]byte, 9)
	row  [N][N]bool    // 每行1～9使用情况
	col  [N][N]bool    // 每列1～9使用情况
	cell [3][3][N]bool // 每个九宫格1～9使用情况
)

func dfs(x, y int) bool {
	defer ot.Flush()
	if y == 9 {
		return dfs(x+1, 0)
	}
	if x == 9 {
		for i := 0; i < 9; i++ {
			fmt.Fprintf(ot, "%s\n", g[i])
		}
		return true
	}

	// 当前不是空的时候搜索下一个
	if g[x][y] != '.' {
		return dfs(x, y+1)
	}

	// 遍历数独中的数字1～9
	for i := 1; i <= 9; i++ {
		if row[x][i] == false && col[y][i] == false && cell[x/3][y/3][i] == false {
			row[x][i], col[y][i], cell[x/3][y/3][i] = true, true, true
			g[x][y] = byte(i + '0')
			if dfs(x, y+1) { // 剪枝
				return true
			}
			g[x][y] = '.'
			row[x][i], col[y][i], cell[x/3][y/3][i] = false, false, false
		}
	}

	return false
}

func main() {
	defer ot.Flush()

	for i := 0; i < 9; i++ {
		fmt.Fscanln(in, &g[i])
		for j := 0; j < 9; j++ {
			if g[i][j] != '.' {
				t := g[i][j] - '0' // 转换坐标
				row[i][t], col[j][t], cell[i/3][j/3][t] = true, true, true
			}
		}
	}

	dfs(0, 0)
}
