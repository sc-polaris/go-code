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

	/*
		// 版本1：
		for i := 1; i <= n; i++ {
			var v, w int
			fmt.Fscan(in, &v, &w)
			for j := 1; j <= m; j++ {
				for k := 0; k*v <= j; k++ {
					f[i][j] = max(f[i][j], f[i-1][j-k*v]+k*w)
				}
			}
		}
		fmt.Fprintln(ot, f[n][m])
	*/

	// 优化1：
	/*
		f[i, j]   = max(f[i-1, j] , f[i-1,j-v]+w ,  f[i-1,j-2*v]+2*w , f[i-1,j-3*v]+3*w , .....)
		f[i, j-v] = max(			f[i-1, j-v]  ,  f[i-1,j-2*v] + w , f[i-1,j-3*v]+2*w , .....)
		由上两式，可得出如下递推关系：
		f[i][j]=max(f[i,j-v]+w , f[i-1][j])
	*/
	/*
		for i := 1; i <= n; i++ {
			var v, w int
			fmt.Fscan(in, &v, &w)
			for j := 0; j <= m; j++ {
				if j < v {
					f[i][j] = f[i-1][j]
				} else {
					// 01背包：f[i][j] = max(f[i-1][j], f[i-1][j-v]+w)
					// 可以理解为01背包只能使用一次所以后面必然是f[i-1][j-v]+w,不能使用当前物品
					// 完全背包可以使用当前物品
					f[i][j] = max(f[i-1][j], f[i][j-v]+w)
				}
			}
		}
		fmt.Fprintln(ot, f[n][m])
	*/

	// 优化2：
	for i := 1; i <= n; i++ {
		var v, w int
		fmt.Fscan(in, &v, &w)
		for j := v; j <= m; j++ {
			f[j] = max(f[j], f[j-v]+w)
		}
	}
	fmt.Fprintln(ot, f[m])
}
