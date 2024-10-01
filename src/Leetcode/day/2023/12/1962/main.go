package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

// Less 最大堆
func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Pop() any           { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) Push(v any)         { h.IntSlice = append(h.IntSlice, v.(int)) }

func minStoneSum(piles []int, k int) int {
	h := &hp{piles}
	heap.Init(h) // 原地堆化
	for ; k > 0 && piles[0] > 0; k-- {
		piles[0] -= piles[0] / 2 // 直接修改堆顶
		heap.Fix(h, 0)
	}
	res := 0
	for _, x := range piles {
		res += x
	}
	return res
}

func main() {

}
