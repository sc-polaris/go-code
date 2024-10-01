package main

import (
	"container/heap"
	"sort"
)

// hp 继承 Len, Less, Swap
type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func magicTower(nums []int) (ans int) {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum < 0 {
		return -1
	}

	flood := 1
	h := &hp{}
	for _, x := range nums {
		if x < 0 {
			heap.Push(h, x)
		}
		flood += x
		if flood < 1 {
			// 这意味着 x < 0, 所以前面必然会把 x 入堆
			// 所以堆必然不是空的，并且堆顶 <= x
			flood -= heap.Pop(h).(int) // 返回
			ans++
		}
	}
	return
}

func main() {

}
