package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
)

/*
func dfs(u int, st []bool) {
	if u == n {
		for i := 0; i < n; i++ {
			if st[i] {
				fmt.Fprintf(ot, "%d ", i+1)
			}
		}
		fmt.Fprintln(ot)
		return
	}

	st[u] = true // 用当前数字
	dfs(u+1, st)
	st[u] = false // 不用当前数字
	dfs(u+1, st)
}
*/

func dfs(u, st int) {
	if u == n {
		for i := 0; i < n; i++ {
			if st>>i&1 == 1 {
				fmt.Fprintf(ot, "%d ", i+1)
			}
		}
		fmt.Fprintln(ot)
		return
	}

	dfs(u+1, st)        // 不用u这个数
	dfs(u+1, st|(1<<u)) // 用u这个数
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	/*
		// 方法1
		st := make([]bool, n)
		dfs(0, st)
	*/

	/*
		// 方法2
		// st 是每一个状态
		for st := 0; st < 1<<n; st++ {
			// 用指针j遍历二进制数st中的每一位
			for j := 0; j < n; j++ {
				if st>>j&1 == 1 {
					fmt.Fprintf(ot, "%d ", j+1)
				}
			}
			fmt.Fprintln(ot)
		}
	*/

	// 方法3
	dfs(0, 0)
}
