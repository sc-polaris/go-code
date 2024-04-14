package main

import "container/list"

type MyHashSet []list.List

func Constructor() MyHashSet {
	return make(MyHashSet, 1000)
}

func (s MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	idx := s.hash(key)
	s[idx].PushFront(key)
}

func (s MyHashSet) Remove(key int) {
	idx := s.hash(key)
	for e := s[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			s[idx].Remove(e)
		}
	}
}

func (s MyHashSet) Contains(key int) bool {
	idx := s.hash(key)
	for e := s[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

func (s MyHashSet) hash(key int) int {
	return key % len(s)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
