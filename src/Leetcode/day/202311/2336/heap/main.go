package main

import "container/heap"

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h *hp) Empty() bool       { return len(*h) == 0 }
func (h *hp) Top() int          { return (*h)[0] }
func (h *hp) Push(val any)      { *h = append(*h, val.(int)) }
func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// limit: 最右边的边界
// hp: 右边界左边的新添加的值
type SmallestInfiniteSet struct {
	h     *hp
	limit int
	mp    map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{h: &hp{}, mp: make(map[int]bool), limit: 1}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.h.Len() > 0 {
		res := heap.Pop(s.h).(int)
		s.mp[res] = false
		return res
	}
	res := s.limit
	s.limit++
	return res
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num >= s.limit || s.mp[num] {
		return
	}
	s.mp[num] = true
	heap.Push(s.h, num)
}

/**
* Your SmallestInfiniteSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.PopSmallest();
* obj.AddBack(num);
 */

func main() {

}
