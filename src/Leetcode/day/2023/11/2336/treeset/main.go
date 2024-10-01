package main

import "github.com/emirpasic/gods/sets/treeset"

type SmallestInfiniteSet struct {
	thres int
	*treeset.Set
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{thres: 1, Set: treeset.NewWithIntComparator()}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.Empty() {
		ans := s.thres
		s.thres++
		return ans
	}
	it := s.Iterator()
	it.Next()
	ans := it.Value().(int)
	s.Remove(ans)
	return ans
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.thres {
		s.Add(num)
	}
}

/**
* Your SmallestInfiniteSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.PopSmallest();
* obj.AddBack(num);
 */
