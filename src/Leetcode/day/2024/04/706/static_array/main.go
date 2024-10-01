package main

type MyHashMap []int

func Constructor() MyHashMap {
	data := make(MyHashMap, 1000010)
	for i := range data {
		data[i] = -1
	}
	return data
}

func (m MyHashMap) Put(key int, value int) {
	m[key] = value
}

func (m MyHashMap) Get(key int) int {
	return m[key]
}

func (m MyHashMap) Remove(key int) {
	m[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
