package main

type NumArray []int

func Constructor(nums []int) NumArray {
	s := make(NumArray, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	return s
}

func (s NumArray) SumRange(left int, right int) int {
	return s[right+1] - s[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {

}
