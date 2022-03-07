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
	a  [N]int
	f  [N]int
	g  [N]int
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 从0开始

func solve2() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// 预处理f[i]：以i结尾的单调上升子串的最大长度
	f[0] = 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}
	}

	// 预处理g[i]：以i开头的单调上升子串的最大长度
	g[n-1] = 1
	for i := n - 1; i >= 0; i-- {
		if a[i] < a[i+1] {
			g[i] = g[i+1] + 1
		} else {
			g[i] = 1
		}
	}

	var res int
	// 枚举删除哪个数
	for i := 0; i < n; i++ {
		if i == 0 {
			res = max(res, g[i+1])
		} else if i == n-1 {
			res = max(res, f[i-1])
		} else if a[i-1] >= a[i+1] {
			res = max(res, max(f[i-1], g[i+1]))
		} else {
			res = max(res, f[i-1]+g[i+1])
		}
	}
	fmt.Fprintln(ot, res)
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// 预处理f[i]：以i结尾的单调上升子串的最大长度
	for i := 1; i <= n; i++ {
		if a[i] > a[i-1] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}
	}

	// 预处理g[i]：以i开头的单调上升子串的最大长度
	for i := n; i > 0; i-- {
		if a[i] < a[i+1] {
			g[i] = g[i+1] + 1
		} else {
			g[i] = 1
		}
	}

	var res int
	// 枚举删除哪个数
	for i := 1; i <= n; i++ {
		if a[i-1] >= a[i+1] {
			res = max(res, max(f[i-1], g[i+1]))
		} else {
			res = max(res, f[i-1]+g[i+1])
		}
	}
	fmt.Fprintln(ot, res)
}
