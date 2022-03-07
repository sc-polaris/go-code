package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)

	mid := a[n/2] // 中位数
	var res int
	for i := range a {
		res += abs(a[i] - mid)
	}

	fmt.Fprintln(ot, res)
}
