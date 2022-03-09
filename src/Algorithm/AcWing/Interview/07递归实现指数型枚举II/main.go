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
)

func dfs(u int) {
	if u == n {
		for i, num := range nums {
			if st[i] {
				fmt.Fprintf(ot, "%d ", num)
			}
		}
		fmt.Fprintln(ot)
		return
	}

	k := u
	for k < n && nums[k] == nums[u] {
		k++
	}

	dfs(k)

	for i := u; i < k; i++ {
		st[i] = true
		dfs(k)
	}

	for i := u; i < k; i++ {
		st[i] = false
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	nums = make([]int, n)
	st = make([]bool, n)

	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	sort.Ints(nums)

	dfs(0)
}
