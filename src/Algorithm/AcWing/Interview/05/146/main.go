package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	Sum int
	Pos int
}

type MinHeap []*Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Sum < h[j].Sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Top() *Node {
	return (*h)[0]
}

func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(*Node))
}

func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

var (
	in      = bufio.NewReader(os.Stdin)
	ot      = bufio.NewWriter(os.Stdout)
	m, n    int
	a, b, c []int // a和b是每次需要合并的序列, c是中间临时缓存的序列
)

func merge() {
	minH := new(MinHeap)

	for i := range b {
		// 后者表示当前a的下标
		heap.Push(minH, &Node{b[i] + a[0], 0})
	}

	for i := range a {
		t := minH.Top()
		heap.Pop(minH)
		s, p := t.Sum, t.Pos
		c[i] = s
		if p+1 < len(a) {
			heap.Push(minH, &Node{s - a[p] + a[p+1], p + 1})
		}
	}

	copy(a, c)
}

func main() {
	defer ot.Flush()

	var T int
	fmt.Fscan(in, &T)

	for ; T > 0; T-- {
		fmt.Fscan(in, &m, &n)
		a, b, c = make([]int, n), make([]int, n), make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		sort.Ints(a)

		for i := 0; i < m-1; i++ {
			for j := range b {
				fmt.Fscan(in, &b[j])
			}
			merge() //每次合并两组, 共合并m-1次
		}

		for i := range a {
			fmt.Fprintf(ot, "%d ", a[i])
		}
		fmt.Fprintln(ot)
	}
}
