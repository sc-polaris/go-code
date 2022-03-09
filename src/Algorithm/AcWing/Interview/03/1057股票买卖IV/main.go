package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 1e5 + 10
	M   = 110
	INF = 0x3f3f3f3f
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, k int
	w    [N]int
	//f    [N][M][2]int // [天数][交易次数][1/0是否持有股票]
	// 股票状态: 奇数表示第 k 次交易持有/买入, 偶数表示第 k 次交易不持有/卖出, 0 表示没有操作
	//f [N][M * 2]int // [天数][股票状态]
	f [M * 2]int // [股票状态]
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &k)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}

	/*  // 三维数组：
	for i := 0; i <= k; i++ {
		f[0][i][1] = -w[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			// dp方程, 0表示不持有/卖出, 1表示持有/买入
			f[i][j][0] = max(f[i-1][j][0], f[i-1][j][1]+w[i])
			f[i][j][1] = max(f[i-1][j][1], f[i-1][j-1][0]-w[i])
		}
	}

	fmt.Fprintln(ot, f[n-1][k][0])
	*/

	/*	// 二维数组：
		// 股票状态: 奇数表示第 k 次交易持有/买入, 偶数表示第 k 次交易不持有/卖出, 0 表示没有操作
		for i := 1; i < k*2; i += 2 {
			f[0][i] = -w[0]
		}

			for i := 1; i < k*2; i += 2 {
			f[0][i] = -w[0]
		}

		for i := 1; i < n; i++ {
			for j := 1; j < k*2; j += 2 {
				f[i][j] = max(f[i-1][j], f[i-1][j-1]-w[i])
				f[i][j+1] = max(f[i-1][j+1], f[i-1][j]+w[i])
			}
		}

		fmt.Fprintln(ot, f[n-1][k*2])
	*/

	// 股票状态: 奇数表示第 k 次交易持有/买入, 偶数表示第 k 次交易不持有/卖出, 0 表示没有操作
	for i := 1; i < k*2; i += 2 {
		f[i] = -w[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < k*2; j += 2 {
			f[j] = max(f[j], f[j-1]-w[i])
			f[j+1] = max(f[j+1], f[j]+w[i])
		}
	}

	fmt.Fprintln(ot, f[k*2])
}
