package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 1e5 + 10
	MOD = 1e9 + 7
)

var (
	in        = bufio.NewReader(os.Stdin)
	ot        = bufio.NewWriter(os.Stdout)
	n1, n2, m int
	f         [N]int
)

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n1, &n2, &m)

	var p int

	f[0] = 1 // 币值为0有一种方案

	// 普通币 完全背包
	for i := 1; i <= n1; i++ {
		fmt.Fscan(in, &p)
		for j := p; j <= m; j++ {
			f[j] = (f[j] + f[j-p]) % MOD
		}
	}

	// 纪念币 01背包
	for i := 1; i <= n2; i++ {
		fmt.Fscan(in, &p)
		for j := m; j >= p; j-- {
			f[j] = (f[j] + f[j-p]) % MOD
		}
	}

	fmt.Fprintln(ot, f[m])
}
