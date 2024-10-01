package main

import "sort"

type SmallestInfiniteSet struct {
	stack    []int // 比small小的 ascend 栈
	smallest int   // 无穷正数，最小值
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{[]int{}, 1}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if len(s.stack) > 0 {
		res := s.stack[0]
		s.stack = s.stack[1:]
		return res
	}
	res := s.smallest
	s.smallest++
	return res
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	// check is num in stack
	if idx := sort.SearchInts(s.stack, num); idx < len(s.stack) && s.stack[idx] == num {
		return
	}
	if s.smallest > num {
		s.stack = append(s.stack, num)
		sort.Ints(s.stack)
	}
}

/**
* Your SmallestInfiniteSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.PopSmallest();
* obj.AddBack(num);
 */

func main() {

}
