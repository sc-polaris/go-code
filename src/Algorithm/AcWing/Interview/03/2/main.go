package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)

	n, m int
	//f    [N][N]int // 前i个物品j容量的最大价值
	f [N]int // 容量为j的最大价值
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	/*	二维数组：
		for i := 1; i <= n; i++ {
			var v, w int
			fmt.Fscan(in, &v, &w)
			for j := 0; j <= m; j++ {
				if j < v {
					f[i][j] = f[i-1][j]
				} else {
					f[i][j] = max(f[i-1][j], f[i-1][j-v]+w)
				}
			}
		}
		fmt.Fprintln(ot, f[n][m])
	*/

	for i := 1; i <= n; i++ {
		var v, w int
		fmt.Fscan(in, &v, &w)
		for j := m; j >= v; j-- {
			f[j] = max(f[j], f[j-v]+w)
		}
	}

	fmt.Fprintln(ot, f[m])

}
