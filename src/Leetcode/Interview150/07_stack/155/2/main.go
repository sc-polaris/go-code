package main

import "math"

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   math.MaxInt,
	}
}

func (s *MinStack) Push(val int) {
	if len(s.stack) == 0 {
		s.min = val
	}
	s.stack = append(s.stack, val-s.min)
	if s.min > val {
		s.min = val
	}
}

func (s *MinStack) Pop() {
	//  如果栈顶元素是负数（比之前的 min 小），则当前栈顶元素就是目前的 min， pop出来之后需要更新 min = min - （栈顶数值）
	if s.stack[len(s.stack)-1] < 0 {
		s.min -= s.stack[len(s.stack)-1]
	}
	// 如果栈顶元素是非负数（比之前的 min 大，或相同），则直接pop栈顶元素
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	// 如果栈顶元素是正数（比之前的 min 大），则 top 数值 = min + （栈顶数值）
	if s.stack[len(s.stack)-1] >= 0 {
		return s.min + s.stack[len(s.stack)-1]
	}
	// 如果栈顶元素是非正数（比之前的 min 小，或相同），则当前 top 元素就等于目前的 min
	return s.min
}

func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
