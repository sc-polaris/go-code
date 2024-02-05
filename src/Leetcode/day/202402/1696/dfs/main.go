package main

import "math"

func maxResult(nums []int, k int) int {
	// dfs(i) 定义为从 0 到 i，经过的所有数字之和的最大值。
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return nums[0]
		}
		mx := math.MinInt
		for j := max(i-k, 0); j < i; j++ {
			mx = max(mx, dfs(j))
		}
		return mx + nums[i]
	}
	return dfs(len(nums) - 1)
}

func main() {

}
