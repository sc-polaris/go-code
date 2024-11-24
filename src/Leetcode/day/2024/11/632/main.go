package main

import (
	"container/heap"
	"math"
	"slices"
)

/*
	你有 k 个 非递减排列 的整数列表。找到一个 最小 区间，使得 k 个列表中的每个列表至少有一个数包含在其中。

	我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。
*/

func smallestRange(nums [][]int) []int {
	h := make(hp, len(nums))
	r := math.MinInt
	for i, arr := range nums {
		h[i] = tuple{arr[0], i, 0} // 把每个列表的第一个元素入堆
		r = max(r, arr[0])
	}
	heap.Init(&h)

	ansL, ansR := h[0].x, r            // 第一个合法区间的左右端点
	for h[0].j+1 < len(nums[h[0].i]) { // 堆顶列表有下一个元素
		x := nums[h[0].i][h[0].j+1] // 堆顶列表的下一个元素
		r = max(r, x)               // 更新合法区间的右端点
		h[0].x = x                  // 替换堆顶
		h[0].j++
		heap.Fix(&h, 0)
		l := h[0].x // 当前合法区间的左端点
		if r-l < ansR-ansL {
			ansL, ansR = l, r
		}
	}
	return []int{ansL, ansR}
}

type tuple struct{ x, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].x < h[j].x }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {} // 没用到，可以不写
func (hp) Pop() (_ any)         { return }

func smallestRange2(nums [][]int) []int {
	type pair struct{ x, i int }
	var pairs []pair
	for i, arr := range nums {
		for _, x := range arr {
			pairs = append(pairs, pair{x, i})
		}
	}
	slices.SortFunc(pairs, func(a, b pair) int { return a.x - b.x })

	ansL, ansR := pairs[0].x, pairs[len(pairs)-1].x
	empty := len(nums)
	cnt := make([]int, empty)
	left := 0
	for _, p := range pairs {
		r, i := p.x, p.i
		if cnt[i] == 0 { // 包含 nums[i] 的数字
			empty--
		}
		cnt[i]++
		for empty == 0 { // 每个列表都至少包含一个数
			l, i := pairs[left].x, pairs[left].i
			if r-l < ansR-ansL {
				ansL, ansR = l, r
			}
			cnt[i]--
			if cnt[i] == 0 {
				// 不包含 nums[i] 的数字
				empty++
			}
			left++
		}
	}
	return []int{ansL, ansR}
}
