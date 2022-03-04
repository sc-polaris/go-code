package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

var (
	n  int
	a  [N]int
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var res int
	for i := 1; i < n; i++ {
		res += max(0, a[i]-a[i-1])
	}

	fmt.Fprintln(ot, res)
}
