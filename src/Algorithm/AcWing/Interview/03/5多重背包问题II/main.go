package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 11010 // // 1 << 11 = 2048 一个s最多分11个包，1000个就是11000
	M = 2010
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	v    [N]int
	w    [N]int
	f    [M]int
	cnt  int
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

	for ; n > 0; n-- {
		var v1, w1, s int
		fmt.Fscan(in, &v1, &w1, &s)
		// 读入s个物品顺便打包
		k := 1
		for k <= s {
			cnt++ // 实际数量
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
