package main

import "math/rand"

/*
	核心思想：
	1. 哈希表记录加入和删除的数，可以O(1)检查是否出现过
	2. 用数组维护所有数，方便随机取一个数，数组后加入一个数也是O(1)，唯一难点在于删除。
	3. 用哈希表维护每个数加入时的坐标，在要删除的数不是数组最后一个时，与最后一个交换（因为是不在乎顺序的，所以这种交换不
	   影响任何东西），此时要删除的数成为数组最后一个，可以O(1)删除
*/

type RandomizedSet struct {
	nums   []int
	idxMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 0), make(map[int]int)}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.idxMap[val]; !ok {
		s.idxMap[val] = len(s.nums)
		s.nums = append(s.nums, val)
		return true
	}
	return false
}

func (s *RandomizedSet) Remove(val int) bool {
	if idx, ok := s.idxMap[val]; ok {
		swapVal := s.nums[len(s.nums)-1]
		s.nums[idx] = swapVal
		s.idxMap[swapVal] = idx
		delete(s.idxMap, val)
		s.nums = s.nums[:len(s.nums)-1]
		return true
	}
	return false
}

func (s *RandomizedSet) GetRandom() int {
	return s.nums[rand.Intn(len(s.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
