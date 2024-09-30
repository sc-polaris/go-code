package main

import (
	"container/heap"
	"sort"
)

/*
	方法二：维护取消预约的座位
	时间复杂度：初始化为 O(1)，reserve 和 unreserve 均为 O(logq)，其中 q 是 unreserve 的调用次数。
	空间复杂度：O(q)。
*/

type SeatManager struct {
	sort.IntSlice
	seats int
}

func Constructor(n int) SeatManager {
	return SeatManager{}
}

func (m *SeatManager) Reserve() int {
	if len(m.IntSlice) > 0 { // 有空出来的椅子
		return heap.Pop(m).(int) // 坐编号最小的
	}
	m.seats += 1 // 添加一把新的椅子
	return m.seats
}

func (m *SeatManager) Unreserve(seatNumber int) {
	heap.Push(m, seatNumber) // 有人离开了椅子
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
