package main

/*
	设 h = nums[p] 是矩形的高度，那么矩形的宽度是多少？我们需要知道：
	1. 在 p 左侧的小于 h 的最近元素的下标 left
	2. 在 p 右侧的小于 h 的最近元素的下包 right
*/

func maximumScore(nums []int, k int) (ans int) {
	n := len(nums)
	left := make([]int, n)
	var stk []int
	for i, x := range nums {
		for len(stk) > 0 && x <= nums[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		} else {
			left[i] = -1
		}
		stk = append(stk, i)
	}

	right := make([]int, n)
	stk = stk[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && nums[i] <= nums[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			right[i] = stk[len(stk)-1]
		} else {
			right[i] = -1
		}
		stk = append(stk, i)
	}

	for i, h := range nums {
		l, r := left[i], right[i]
		if l < k && k < r {
			ans = max(ans, h*(r-l-1))
		}
	}
	return
}

func main() {

}
