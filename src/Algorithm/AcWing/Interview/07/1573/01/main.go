package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	nums []int
	st   []bool
)

// u:数的深度，start：开始的数
func dfs(u, start int) {
	if u == m {
		for i := 0; i < n; i++ {
			if st[i] {
				fmt.Fprintf(ot, "%d ", nums[i])
			}
		}
		fmt.Fprintln(ot)
		return
	}

	// 构建几个分支
	for i := start; i < n; i++ {
		if i != 0 && !st[i-1] && nums[i-1] == nums[i] {
			continue
		}
		st[i] = true
		dfs(u+1, i+1)
		st[i] = false
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	nums = make([]int, n)
	st = make([]bool, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	sort.Ints(nums)

	dfs(0, 0)
}
