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

func dfs(u, s int) {
	if s > m {
		return
	}
	if s == m {
		for i := 0; i < u; i++ {
			if st[i] {
				fmt.Fprintf(ot, "%d ", nums[i])
			}
		}
		fmt.Fprintln(ot)
		return
	}

	if u == n {
		return
	}

	var k int
	for k = u; k < n && nums[k] == nums[u]; k, s = k+1, s+1 {
		st[k] = true
	}

	dfs(k, s) // 当前这段全选
	for i := k - 1; i >= u; i-- {
		st[i] = false
		s--
		dfs(k, s)
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
