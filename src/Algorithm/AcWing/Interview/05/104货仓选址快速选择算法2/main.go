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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func swap(q []int, i, j int) {
	q[i], q[j] = q[j], q[i]
}

func quickSelect(q []int, l, r, k int) int {
	if l >= r {
		return q[k]
	}

	i, j, x := l-1, r+1, q[(l+r)>>1]
	for i < j {
		for i = i + 1; q[i] < x; i++ {
		}
		for j = j - 1; q[j] > x; j-- {
		}
		if i < j {
			swap(q, i, j)
		}
	}

	if j >= k {
		return quickSelect(q, l, j, k)
	} else {
		return quickSelect(q, j+1, r, k)
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var mid int
	// 奇数直接取中位数
	if n&1 == 1 {
		mid = quickSelect(a, 0, n-1, (n-1)>>1)
	} else { // 偶数取中间两个数的和
		mid = (quickSelect(a, 0, n-1, n>>1) + quickSelect(a, 0, n-1, (n-1)>>1)) >> 1
	}

	var res int
	for i := range a {
		res += abs(a[i] - mid)
	}

	fmt.Fprintln(ot, res)
}
