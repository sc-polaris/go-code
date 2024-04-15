package main

import "container/list"

const base = 769

type entry struct{ key, value int }

type MyHashMap []list.List

func Constructor() MyHashMap {
	return make(MyHashMap, base)
}

func (m MyHashMap) hash(key int) int {
	return key % base
}

func (m MyHashMap) Put(key int, value int) {
	idx := m.hash(key)
	for e := m[idx].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m[idx].PushBack(entry{key, value})
}

func (m MyHashMap) Get(key int) int {
	idx := m.hash(key)
	for e := m[idx].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m MyHashMap) Remove(key int) {
	idx := m.hash(key)
	for e := m[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m[idx].Remove(e)
		}
	}
}
