package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

/*
heap是常用的实现优先队列的方法。heap包对任意实现了heap接口的类型提供堆操作。
堆结构继承自sort.Interface, 而sort.Interface，需要实现三个方法：
Len() int / Less(i, j int) bool / Swap(i, j int) 再加上堆接口定义的两个方法：
Push(x interface{}) / Pop() interface{}。故只要实现了这五个方法，便定义了一个堆。
*/

// 小根堆

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Empty() bool {
	return len(*h) == 0
}

func (h *minHeap) Top() int {
	return (*h)[0]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 删除堆尾的元素，注意和heap.Pop()区分
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var (
	in = bufio.NewReader(os.Stdin)
	ot = bufio.NewWriter(os.Stdout)
)

func main() {
	h := &maxHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", h.Top())
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}
