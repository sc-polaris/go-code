package list

import "container/list"

type MinStack struct {
	stackValue *list.List
	stackMin   *list.List
}

/** initialize your data structure here. */

func Constructor() MinStack {
	return MinStack{
		stackValue: list.New(),
		stackMin:   list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.stackValue.PushBack(x)
	if this.stackMin.Len() == 0 || this.stackMin.Back().Value.(int) >= x {
		this.stackMin.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	value := this.stackValue.Back()
	minValue := this.stackMin.Back()
	if minValue.Value.(int) == value.Value.(int) {
		this.stackMin.Remove(minValue)
	}
	this.stackValue.Remove(value)
}

func (this *MinStack) Top() int {
	return this.stackValue.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.stackMin.Back().Value.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
