package main

/*
	给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
*/

func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for p := n - 1; p >= 0; p-- {
		if x, y := nums[i], nums[j]; -x > y {
			ans[p] = x * x
			i++
		} else {
			ans[p] = y * y
			j--
		}
	}
	return ans
}
