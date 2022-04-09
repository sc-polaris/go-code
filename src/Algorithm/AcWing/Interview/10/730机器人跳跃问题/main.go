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
	h  []int
)

func check(E int) bool {
	for _, v := range h {
		//if v > E {
		//	E -= v - E
		//} else {
		//	E += E - v
		//}
		// 即：
		E = E*2 - v
		if E > 100000 {
			return true
		}
		if E < 0 {
			return false
		}
	}

	return true
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	h = make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}

	l, r := 0, 100000
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	fmt.Fprintln(ot, l)
}
