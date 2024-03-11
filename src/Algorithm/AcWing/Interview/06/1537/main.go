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
	n    int
	nums []int
	st   []bool
	res  []int
)

func dfs(u int) {
	if u == n {
		for _, v := range res {
			fmt.Fprintf(ot, "%d ", v)
		}
		fmt.Fprintln(ot)
		return
	}

	for i := 0; i < n; i++ {
		if !st[i] {
			res[u] = nums[i]
			st[i] = true
			dfs(u + 1)
			st[i] = false
			for i+1 < n && nums[i+1] == nums[i] {
				i++
			}
		}
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	res, nums = make([]int, n), make([]int, n)
	st = make([]bool, n)

	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	sort.Ints(nums)

	dfs(0)
}
