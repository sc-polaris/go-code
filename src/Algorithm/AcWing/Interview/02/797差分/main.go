package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

var n, m int
var a [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := n; i > 0; i-- {
		a[i] -= a[i-1]
	}

	for ; m > 0; m-- {
		var l, r, c int
		fmt.Fscan(in, &l, &r, &c)
		a[l] += c
		a[r+1] -= c
	}

	for i := 1; i <= n; i++ {
		a[i] += a[i-1]
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
