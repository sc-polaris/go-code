package main

import (
	"container/heap"
	"sort"
)

/*
	给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

	一次操作中，你将执行：
	· 选择 nums 中最小的两个整数 x 和 y 。
	· 将 x 和 y 从 nums 中删除。
	· 将 min(x, y) * 2 + max(x, y) 添加到数组中的任意位置。
	注意，只有当 nums 至少包含两个元素时，你才可以执行以上操作。

	你需要使数组中的所有元素都大于或等于 k ，请你返回需要的 最少 操作次数。
*/

func minOperations(nums []int, k int) (ans int) {
	h := &hp{nums}
	heap.Init(h)
	for h.IntSlice[0] < k {
		x := heap.Pop(h).(int)
		h.IntSlice[0] += x * 2
		heap.Fix(h, 0)
		ans++
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(any) {}
func (h *hp) Pop() any { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
