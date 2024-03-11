package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// 小根堆

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Empty() bool {
	return h.Len() == 0
}

func (h *minHeap) Top() int {
	return (*h)[0]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 大根堆

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Empty() bool {
	return h.Len() == 0
}

func (h *maxHeap) Top() int {
	return (*h)[0]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func main() {
	defer ot.Flush()

	var T int //T为数据集个数
	fmt.Fscan(in, &T)

	for ; T > 0; T-- {
		var id, n int
		fmt.Fscan(in, &id, &n)
		// 输出数据集编号和中位数个数(即奇数位个数)
		fmt.Fprintf(ot, "%d %d\n", id, (n+1)>>1)

		down, up := new(maxHeap), new(minHeap) // 大根堆、小根堆

		var cnt int // 用于分隔输出,每十个数一行
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(in, &x)

			// 下面为空或x小于下方堆顶,则将x插入大根堆
			if down.Empty() || x <= down.Top() {
				heap.Push(down, x)
			} else {
				heap.Push(up, x)
			}

			// 如果有偶数个数，上面和下面一样多，如果有奇数个数，则下面比上面多一个
			// 如果下面比上面多超过1个
			if down.Len()-up.Len() >= 2 {
				heap.Push(up, down.Top())
				heap.Pop(down)
			}
			if up.Len() > down.Len() {
				// 上面多了挤一个放下面
				heap.Push(down, up.Top())
				heap.Pop(up)
			}

			if i%2 == 0 {
				fmt.Fprintf(ot, "%d ", down.Top())
				cnt++
				if cnt%10 == 0 {
					fmt.Fprintln(ot)
				}
			}
		}

		if cnt%10 != 0 {
			fmt.Fprintln(ot)
		}
	}
}
