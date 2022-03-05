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
	//f  [N][3]int // [天数][操作状态]
	f [3]int
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

	/*	二维数组：
		// f[i][0]: 不持股且当天没卖出
		// f[i][1]: 持股
		// f[i][2]: 不持股且当天卖出
		f[0][0], f[0][1], f[0][2] = 0, -w[0], 0
		for i := 1; i < n; i++ {
			f[i][0] = max(f[i-1][0], f[i-1][2])
			f[i][1] = max(f[i-1][1], f[i-1][0]-w[i])
			f[i][2] = f[i-1][1] + w[i]
		}

		fmt.Fprintln(ot, max(f[n-1][0], f[n-1][2]))
	*/
	// f[0]: 不持股且当天没卖出
	// f[1]: 持股
	// f[2]: 不持股且当天卖出
	f[0], f[1], f[2] = 0, -w[0], 0
	for i := 1; i < n; i++ {
		new0 := max(f[0], f[2])
		new1 := max(f[1], f[0]-w[i])
		new2 := f[1] + w[i]
		f[0], f[1], f[2] = new0, new1, new2
	}
	fmt.Fprintln(ot, max(f[0], f[2]))
}
