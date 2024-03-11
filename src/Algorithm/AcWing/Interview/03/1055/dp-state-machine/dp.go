package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
dp-state-machine：
	f[i][j]：
		集合：考虑前i天的股市，第j天手中持股状态为j（j=0未持股、j=1持股）的方案
		属性：方案的收益最大值
1. 当前处于未持股状态0：
对应可以进行的转换：
    0->0 （不买入，继续观望，那么就什么都不发生）
    0->1 （买入股票，那么收益就要减去当前市场的股票价格）
2. 当前处于持股状态1：
对应可以进行的转换：
    1->1 （不卖出，继续观望，那么就什么都不发生）
    1->0 （卖出股票，那么收益就要加上当前市场的股票价格）
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
	f  [N][2]int
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

	// 第一天: 未持股/持股
	f[0][0], f[0][1] = 0, -w[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]+w[i])
		f[i][1] = max(f[i-1][1], f[i-1][0]-w[i])
	}

	fmt.Fprintln(ot, f[n-1][0])
	/*	优化一维：
		// 第1天: 未持股/持股
		f[0], f[1] := 0, -w[0]
		for i := 1; i < n; i++ {
			f[0], f[1] = max(f[0], f[1]+w[i]), max(f[1], f0-w[i])
		}

		fmt.Fprintln(ot, f[0])
	*/
}
