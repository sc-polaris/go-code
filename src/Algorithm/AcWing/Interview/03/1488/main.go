package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 1e5 + 10
	M = 3 * N
)

// 对于每个点k，开一个单链表，存储k所有可以走到的点。h[k]存储这个单链表的头结点
// h[k]存储链表头，e[]存储节点的值，ne[]存储节点的next指针，idx表示当前用到了哪个节点
var (
	in       = bufio.NewReader(os.Stdin)
	ot       = bufio.NewWriter(os.Stdout)
	n, m     int
	h        = make([]int, N)
	e, ne, w [M]int
	idx      int
	dist     = make([]int, N)
	st       [N]bool
)

// add 添加一条边a->b
func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func memset(a []int, v int) {
	for i := range a {
		a[i] = v
	}
}

func spfa() {
	memset(dist, 0x3f3f3f3f)
	//for i := 0; i <= n; i++ {
	//	dist[i] = 0x3f3f3f3f
	//}
	var q []int // 队列
	dist[0] = 0
	st[0] = true
	q = append(q, 0)

	for len(q) > 0 {
		t := q[0] // 队头
		q = q[1:] // 队头出队
		st[t] = false
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
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

	fmt.Fscan(in, &n, &m)

	memset(h, -1)
	// 初始化头节点
	//for i := 0; i <= n; i++ {
	//	h[i] = -1
	//}

	for ; m > 0; m-- {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		add(a, b, c)
		add(b, a, c)
	}

	fmt.Fscan(in, &m)
	for ; m > 0; m-- {
		var v int // 有商店的村庄编号
		fmt.Fscan(in, &v)
		add(0, v, 0) // 和超级源点相连
	}

	spfa()

	fmt.Fscan(in, &m)
	for ; m > 0; m-- {
		var v int
		fmt.Fscan(in, &v)
		fmt.Fprintln(ot, dist[v])
	}
}
