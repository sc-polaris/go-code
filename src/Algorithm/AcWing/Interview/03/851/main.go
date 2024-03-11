package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 1e5 + 10
	M   = 2 * N
	INF = 0x3f3f3f3f
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	n, m int
	h    = make([]int, N)
	e    = make([]int, M)
	ne   = make([]int, M)
	w    = make([]int, M)
	idx  int
	dist = make([]int, N)
	q    = make([]int, N)
	st   = make([]bool, N)
)

func add(a, b, c int) {
	e[idx], w[idx], ne[idx], h[a], idx = b, c, h[a], idx, idx+1
}

func memset(a []int, v int) {
	for i := range a {
		a[i] = v
	}
}

func spfa() {
	memset(dist, INF)
	dist[1] = 0
	st[1] = true
	q = append(q, 1)

	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		// 从队列中取出来之后该节点st被标记为false,代表之后该节点如果发生更新可再次入队
		st[t] = false
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
				// 当前已经加入队列的结点，无需再次加入队列，即便发生了更新也只用更新数值即可，重复添加降低效率
				if !st[j] {
					st[j] = true
					q = append(q, j)
				}
			}
		}
	}
}

func main() {
	defer ot.Flush()

	memset(h, -1)

	fmt.Fscan(in, &n, &m)

	for ; m > 0; m-- {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		add(a, b, c)
	}

	spfa()

	if dist[n] == INF {
		fmt.Fprintln(ot, "impossible")
	} else {
		fmt.Fprintln(ot, dist[n])
	}
}
