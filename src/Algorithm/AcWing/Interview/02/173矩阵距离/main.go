package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 1010
	M = N * N
)

var n, m int
var g [N][]byte
var dist [N][N]int

func bfs() {
	var q [][]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == '1' {
				dist[i][j] = 0
				q = append(q, []int{i, j})
			} else {
				dist[i][j] = -1
			}
		}
	}

	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < n && b >= 0 && b < m && dist[a][b] == -1 {
				dist[a][b] = dist[x][y] + 1
				q = append(q, []int{a, b})
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(in, &n, &m)

	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &g[i])
	}

	bfs()

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(out, "%d ", dist[i][j])
		}
		fmt.Fprintln(out)
	}

	out.Flush()
}
