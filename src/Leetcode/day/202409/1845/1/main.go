package main

import (
	"container/heap"
	"sort"
)

/*
	请你设计一个管理 n 个座位预约的系统，座位编号从 1 到 n 。

	请你实现 SeatManager 类：
	1. SeatManager(int n) 初始化一个 SeatManager 对象，它管理从 1 到 n 编号的 n 个座位。所有座位初始都是可预约的。
	2. int reserve() 返回可以预约座位的 最小编号 ，此座位变为不可预约。
	3. void unreserve(int seatNumber) 将给定编号 seatNumber 对应的座位变成可以预约。
*/

// 方法一：维护可预约的座位
// 时间复杂度：初始化为 O(n) 或 O(nlogn)，取决于实现。reserve 和 unreserve 均为 O(logn)。
// 空间复杂度：O(n)。

type SeatManager struct {
	sort.IntSlice
}

func Constructor(n int) SeatManager {
	m := SeatManager{make(sort.IntSlice, n)}
	for i := range m.IntSlice {
		m.IntSlice[i] = i + 1
	}

	// 有序数组无需堆化
	return m
}

func (m *SeatManager) Reserve() int {
	return heap.Pop(m).(int)
}

func (m *SeatManager) Unreserve(seatNumber int) {
	heap.Push(m, seatNumber)
}

func (m *SeatManager) Push(v any) { m.IntSlice = append(m.IntSlice, v.(int)) }
func (m *SeatManager) Pop() any {
	a := m.IntSlice
	v := a[len(a)-1]
	m.IntSlice = a[:len(a)-1]
	return v
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
