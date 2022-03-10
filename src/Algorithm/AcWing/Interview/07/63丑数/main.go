package main

import "container/heap"

var factors = []int{2, 3, 5}

type MinHeap []int

func (h MinHeap) Len() int              { return len(h) }
func (h MinHeap) Less(i, j int) bool    { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v interface{})   { *h = append(*h, v.(int)) }
func (h *MinHeap) Pop() (v interface{}) { *h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]; return }

func getUglyNumber(n int) int {
	h := &MinHeap{1}
	var m map[int]bool // 模拟C++：set
	m = make(map[int]bool)
	m[1] = true
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, flag := m[next]; !flag {
				heap.Push(h, next)
				m[next] = true
			}
		}
	}
}

// 方法2
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getUglyNumber2(n int) int {
	q := []int{1}
	i, j, k := 0, 0, 0
	for len(q) < n {
		t := min(q[i]*2, min(q[j]*3, q[k]*5))
		q = append(q, t)
		if q[i]*2 == t {
			i++
		}
		if q[j]*3 == t {
			j++
		}
		if q[k]*5 == t {
			k++
		}
	}
	return q[len(q)-1]
}
