package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
)

func dfs(u, s, st int) {
	if s == m {
		for i := 0; i < n; i++ {
			if st>>i&1 == 1 {
				fmt.Fprintf(ot, "%d ", i+1)
			}
		}
		fmt.Fprintln(ot)
		return
	}

	if u == n {
		return
	}

	dfs(u+1, s+1, st|(1<<u))
	dfs(u+1, s, st)
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)

	dfs(0, 0, 0)
}
