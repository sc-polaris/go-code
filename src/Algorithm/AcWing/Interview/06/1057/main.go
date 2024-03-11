package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
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

	var n, k int
	fmt.Fscan(in, &n, &k)

	f := make([]int, 2*k+1)
	w := make([]int, n)

	for i := range w {
		fmt.Fscan(in, &w[i])
	}

	for i := 1; i < 2*k; i += 2 {
		f[i] = -w[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < 2*k; j += 2 {
			f[j] = max(f[j], f[j-1]-w[i])
			f[j+1] = max(f[j+1], f[j]+w[i])
		}
	}

	fmt.Fprintln(ot, f[2*k])

}
