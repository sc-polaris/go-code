package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
)

type Node struct {
	Time   int
	Profit int
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Top() int           { return (*h)[0] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
	n  int
)

func main() {
	defer ot.Flush()

	for {
		_, err := fmt.Fscan(in, &n)
		if err == io.EOF {
			break
		}
		products := make([]Node, n)
		for i := range products {
			fmt.Fscan(in, &products[i].Profit, &products[i].Time)
		}

		sort.Slice(products, func(i, j int) bool {
			return products[i].Time < products[j].Time
		})

		minH := new(MinHeap)
		for _, product := range products {
			heap.Push(minH, product.Profit)
			if minH.Len() > product.Time {
				heap.Pop(minH)
			}
		}

		var res int
		for minH.Len() > 0 {
			res += minH.Top()
			heap.Pop(minH)
		}

		fmt.Fprintln(ot, res)
	}
}
