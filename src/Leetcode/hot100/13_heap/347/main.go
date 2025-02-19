package main

import "container/heap"

func topKFrequent(nums []int, k int) (ans []int) {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}

	h := hp{}
	heap.Init(&h)
	for key, value := range m {
		heap.Push(&h, pair{key, value})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	for range k {
		ans = append(ans, heap.Pop(&h).(pair).val)
	}
	return
}

type pair struct{ val, cnt int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].cnt < h[j].cnt }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func topKFrequent2(nums []int, k int) (ans []int) {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}

	buckets := make([][]int, len(nums)+1)
	for key, value := range m {
		buckets[value] = append(buckets[value], key)
	}

	for i := len(buckets) - 1; k > 0; i-- {
		for _, v := range buckets[i] {
			ans = append(ans, v)
			k--
		}
	}
	return
}
