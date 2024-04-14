package static_array

type MyHashSet []bool

func Constructor() MyHashSet {
	return make(MyHashSet, 1000010)
}

func (s MyHashSet) Add(key int) {
	s[key] = true
}

func (s MyHashSet) Remove(key int) {
	s[key] = false
}

func (s MyHashSet) Contains(key int) bool {
	return s[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
