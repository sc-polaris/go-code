package main

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	ans := make([][]int, 0, k)
	h := hp{{nums1[0] + nums2[0], 0, 0}}
	for len(ans) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j == 0 && i+1 < n {
			heap.Push(&h, tuple{nums1[i+1] + nums2[0], i + 1, 0})
		}
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

type tuple struct{ s, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
