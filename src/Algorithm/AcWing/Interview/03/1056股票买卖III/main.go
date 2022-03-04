package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
dp状态机：
	f[i][j][k]：
		集合：考虑前i天的股市，第j次交易，手中持股状态为k（k=0未持股、k=1持股）的方案
		属性：方案的收益最大值
	f[i][j][0]:第i天，第j次交易，手中无货
	f[i][j][1]:第i天，第j次交易，手中有货

	f[i][j][0] = max(f[i-1][j][0], f[i-1][j][1]+w[i])
	f[i][j][1] = max(f[i-1][j][1], f[i-1][j][0]-w[i])
	f[i][0][0] = 0
*/

const (
	N   = 1e5 + 10
	INF = 0x3f3f3f3f
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
	w  [N]int
	f  [N][3][2]int
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}

	f[0][0][0] = 0     // 第一天休息
	f[0][0][1] = -w[0] // 第一天买入
	f[0][1][0] = -INF  // 第一天不可能已经有卖出
	f[0][1][1] = -INF
	f[0][2][0] = -INF
	f[0][2][1] = -INF
	for i := 1; i < n; i++ {
		f[i][0][0] = 0                                    // 第i+1天休息
		f[i][0][1] = max(f[i-1][0][1], f[i-1][0][0]-w[i]) // 第i+1天买入
		f[i][1][0] = max(f[i-1][1][0], f[i-1][0][1]+w[i]) // 第i+1天第1次卖出
		f[i][1][1] = max(f[i-1][1][1], f[i-1][1][0]-w[i]) // 第i+1天第1次买入
		f[i][2][0] = max(f[i-1][2][0], f[i-1][1][1]+w[i]) // 第i+1天第2次卖出
		f[i][2][1] = max(f[i-1][2][1], f[i-1][2][1]-w[i]) // 第i+1天第2次买入
	}

	fmt.Fprintln(ot, max(0, max(f[n-1][1][0], f[n-1][2][0])))
	/*  优化：
	// 第一次 第二次交易
	buy1, sell1 := -w[0], 0
	buy2, sell2 := -w[0], 0
	for i := 1; i < n; i++ {
		buy1 = max(buy1, 0-w[i])
		sell1 = max(sell1, buy1+w[i])
		buy2 = max(buy2, sell1-w[i])
		sell2 = max(sell2, buy2+w[i])
	}

	fmt.Fprintln(ot, sell2)
	*/
}
