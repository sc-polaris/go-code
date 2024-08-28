package main

/*
	answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积。换句话说，
	如果知道了 i 左边所有数的乘积，以及 i 右边所有数的乘积，就可以算出 answer[i]。

	于是：
	1. 定义 pre[i] 表示从 nums[0] 到 nums[i−1] 的乘积。
	2. 定义 suf[i] 表示从 nums[i+1] 到 nums[n−1] 的乘积。

	我们可以先计算出从 nums[0] 到 nums[i−2] 的乘积 pre[i−1]，再乘上 nums[i−1]，就得到了 pre[i]，即
				pre[i]=pre[i−1]⋅nums[i−1]
	同理有
				suf[i]=suf[i+1]⋅nums[i+1]
	初始值：pre[0]=suf[n−1]=1。
	算出 pre 数组和 suf 数组后，有
				answer[i]=pre[i]⋅suf[i]
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = pre[i] * suf[i]
	}
	return ans
}

// 优化 不使用额外空间
// 先计算 suf，然后一边计算 pre，一边把 pre 直接乘到 suf[i] 中。最后返回 suf。

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	pre := 1
	for i, x := range nums {
		// 此时 pre 为 nums[0] 到 nums[i-1] 的乘积，直接乘到 suf[i] 中
		suf[i] *= pre
		pre *= x
	}
	return suf
}
