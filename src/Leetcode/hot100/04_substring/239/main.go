package main

func maxSlidingWindow(nums []int, k int) (ans []int) {
	var q []int
	for i, x := range nums {
		// 入
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // 维护 q 的单调性
		}
		q = append(q, i)
		// 出
		if i-q[0] >= k { // 队首已经离开窗口了
			q = q[1:]
		}
		// 记录答案
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return
}
