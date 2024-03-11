package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	x int
	y int
}

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	g    [][]byte
	dist [][]int
)

func bfs() {
	var q []Node

	dist = make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if g[i][j] == '1' {
				dist[i][j] = 0
				q = append(q, Node{x: i, y: j})
			} else {
				dist[i][j] = -1
			}
		}
	}

	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		x, y := t.x, t.y
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < n && b >= 0 && b < m && dist[a][b] == -1 {
				dist[a][b] = dist[x][y] + 1
				q = append(q, Node{x: a, y: b})
			}
		}
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n, &m)
	g = make([][]byte, n)

	for i := range g {
		fmt.Fscan(in, &g[i])
	}

	bfs()

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(ot, "%d ", dist[i][j])
		}
		fmt.Fprintln(ot)
	}

}
