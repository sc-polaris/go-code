package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 110

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	f    [N]int
	//f    [N][N]int
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
		// 二维：
		for i := 1; i <= n; i++ {
			var v, w, s int
			fmt.Fscan(in, &v, &w, &s)
			for j := 0; j <= m; j++ {
				for k := 0; k <= s && v*k <= j; k++ {
					f[i][j] = max(f[i][j], f[i-1][j-v*k]+w*k)
				}
			}
		}

		fmt.Fprintln(ot, f[n][m])
	*/

	/*
		// 一维
		for i := 1; i <= n; i++ {
			var v, w, s int
			fmt.Fscan(in, &v, &w, &s)
			for j := m; j >= v; j-- {
				for k := 0; k <= s && v*k <= j; k++ {
					f[j] = max(f[j], f[j-v*k]+w*k)
				}
			}
		}

		fmt.Fprintln(ot, f[m])
	*/

	// 一维化简为01背包问题 这里最多100个s,一个s最大为100，2^7=128>100,所以最多700个物品,N要大于700
	var cnt int
	var v, w [N]int
	for ; n > 0; n-- {
		var v1, w1, s int
		fmt.Fscan(in, &v1, &w1, &s)
		// 读入s个物品顺便打包
		k := 1
		for k <= s {
			cnt++ // 实际物品种数
			v[cnt] = v1 * k
			w[cnt] = w1 * k
			s -= k
			k <<= 1 // 倍增包裹大小
		}

		if s > 0 { // 不足的单独放一个
			cnt++
			v[cnt] = v1 * s
			w[cnt] = w1 * s
		}
	}

	n = cnt // 更新物品种数
	// 01背包模版
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Fprintln(ot, f[m])
}
