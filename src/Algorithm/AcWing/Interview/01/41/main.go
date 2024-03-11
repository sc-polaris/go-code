package main

type MinStack struct {
	stackValue []int
	stackMin   []int
}

/** initialize your data structure here. */

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stackValue = append(this.stackValue, x)
	if len(this.stackMin) == 0 || this.stackMin[len(this.stackMin)-1] >= x {
		this.stackMin = append(this.stackMin, x)
	}

}

func (this *MinStack) Pop() {
	if this.stackMin[len(this.stackMin)-1] == this.stackValue[len(this.stackValue)-1] {
		this.stackMin = this.stackMin[:len(this.stackMin)-1]
	}
	this.stackValue = this.stackValue[:len(this.stackValue)-1]
}

func (this *MinStack) Top() int {
	return this.stackValue[len(this.stackValue)-1]
}

func (this *MinStack) GetMin() int {
	return this.stackMin[len(this.stackMin)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
