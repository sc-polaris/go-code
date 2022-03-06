package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 1010

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
)

func judge(a, b []int, x, y int) int {
	if a[x] > b[y] {
		return 1
	} else if a[x] < b[y] {
		return -1
	} else {
		return 0
	}
}

func main() {
	defer ot.Flush()

	for {
		fmt.Fscan(in, &n)

		if n == 0 {
			return
		}

		a := make([]int, n) // 田忌
		b := make([]int, n) // 国王

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}

		// 从大到小排序
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		sort.Sort(sort.Reverse(sort.IntSlice(b)))
		f1, f2, s1, s2 := 0, 0, n-1, n-1
		res := 0
		for f1 <= s1 {
			if judge(a, b, s1, s2) > 0 {
				res, s1, s2 = res+1, s1-1, s2-1
			} else if judge(a, b, s1, s2) < 0 {
				res, s1, f2 = res-1, s1-1, f2+1
			} else {
				if judge(a, b, f1, f2) > 0 {
					res, f1, f2 = res+1, f1+1, f2+1
				} else {
					res, s1, f2 = res+judge(a, b, s1, f2), s1-1, f2+1
				}
			}
		}

		fmt.Fprintln(ot, res*200)
	}
}
