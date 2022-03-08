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

func main() {
	defer ot.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)

	var sum int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	avg := sum / n
	var res int
	for i, x := 0, 0; i < n; i++ {
		x = a[i] - avg + x
		if x != 0 {
			res++
		}
	}

	fmt.Fprintln(ot, res)
}
