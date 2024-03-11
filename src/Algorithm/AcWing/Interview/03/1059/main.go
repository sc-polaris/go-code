package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
	w  [N]int
	f  int
	//dp [N][2]int
	dp [2]int
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &f)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}

	/*	二维数组：
		// dp[i][j]:第i天状态j(0/1卖出买入)
		dp[0][0], dp[0][1] = 0, -w[0]
		for i := 1; i < n; i++ {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+w[i]-f)
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-w[i])
		}

		fmt.Fprintln(ot, dp[n-1][0])
	*/

	// dp[j]:状态j(0/1卖出买入)
	dp[0], dp[1] = 0, -w[0]
	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], dp[1]+w[i]-f)
		dp[1] = max(dp[1], dp[0]-w[i])
	}
	fmt.Fprintln(ot, dp[0])
}
