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
	//f  [N][3][2]int // 最多2次交易，状态0/1
	//f [N][5]int
	f [5]int
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

	/*	最细：
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
			f[i][2][1] = -INF 								  // 第i+1天第2次不可能买入
		}

		fmt.Fprintln(ot, max(0, max(f[n-1][1][0], f[n-1][2][0])))
	*/

	// f[i][j]:第i天，j为 [0 - 4] 五个状态，f[i][j]表示第i天状态j所剩最大现金。
	/* 二维数组写法：
	 * 定义 5 种状态:
	 * 0: 没有操作, 1: 第一次买入, 2: 第一次卖出, 3: 第二次买入, 4: 第二次卖出
	 */

	/*
		f[0][1], f[0][3] = -w[0], -w[0]
		for i := 1; i < n; i++ {
			f[i][0] = f[i-1][0]
			f[i][1] = max(f[i-1][1], f[i-1][0]-w[i])
			f[i][2] = max(f[i-1][2], f[i-1][1]+w[i])
			f[i][3] = max(f[i-1][3], f[i-1][2]-w[i])
			f[i][4] = max(f[i-1][4], f[i-1][3]+w[i])
		}
		fmt.Fprintln(ot, f[n-1][4])
	*/

	// 优化：一维
	f[1], f[2] = -w[0], 0
	f[3], f[4] = -w[0], 0
	for i := 1; i < n; i++ {
		f[1] = max(f[1], f[0]-w[i])
		f[2] = max(f[2], f[1]+w[i])
		f[3] = max(f[3], f[2]-w[i])
		f[4] = max(f[4], f[3]+w[i])
	}

	fmt.Fprintln(ot, f[4])

}
