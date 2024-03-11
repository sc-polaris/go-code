package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 20
	M = 1 << 20
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
	f  [M][N]int // f[i][j]表示所有从0走到j，走过的所有的点的情况是i的所有路径
	w  [N][N]int
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &w[i][j])
		}
	}

	// 初始化f数组
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = 0x3f3f3f3f
		}
	}

	f[1][0] = 0                 //因为0是起点，所以f[1][0]=0
	for i := 0; i < 1<<n; i++ { // i 表示所有的情况
		for j := 0; j < n; j++ { // j表示走到哪一个点
			if i>>j&1 == 1 { // 查看第j个点是否走过
				for k := 0; k < n; k++ { // k表示走到j这个点之前，以k为终点的最短距离
					if i>>k&1 == 1 {
						f[i][j] = min(f[i][j], f[i-(1<<j)][k]+w[k][j]) // 更新最短距离
					}
				}
			}
		}
	}

	fmt.Fprintln(ot, f[(1<<n)-1][n-1])

}
