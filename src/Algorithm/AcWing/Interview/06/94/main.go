package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n    int
	nums []int
	st   []bool
)

func dfs(u int) {
	if u == n {
		for i := range nums {
			fmt.Fprintf(ot, "%d ", nums[i]+1)
		}
		fmt.Fprintln(ot)
		return
	}

	for i := 0; i < n; i++ {
		if !st[i] {
			nums[u] = i
			st[i] = true
			dfs(u + 1)
			st[i] = false
		}
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	nums = make([]int, n)
	st = make([]bool, n)

	dfs(0)
}
