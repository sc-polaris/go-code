package main

func maximumSum(nums []int) int {
	ans := -1
	maxV := make([]int, 82) // 最多9个9相加
	for _, v := range nums {
		s := 0 // 数位和
		for x := v; x > 0; x /= 10 {
			s += x % 10
		}
		if maxV[s] > 0 { // 说明左边也有数位和等于s的元素
			ans = max(ans, maxV[s]+v) // 更新最大值
		}
		maxV[s] = max(maxV[s], v) // 维护数位和最大的
	}
	return ans
}
