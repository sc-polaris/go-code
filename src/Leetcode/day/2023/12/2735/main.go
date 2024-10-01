package main

import "slices"

func minCost(nums []int, x int) int64 {
	n := len(nums)
	s := make([]int64, n) // s[k] 统计操作 k 次的总成本
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums { // 子数组的左断点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形
			mn = min(mn, nums[j%n]) // 维护从 nums[i] 到 nums[j] 的最小值
			s[j-i] += int64(mn)     // 累加操作 j-i 次的花费
		}
	}
	return slices.Min(s)
}

func minCost2(nums []int, x int) int64 {
	sum := func(arr []int) (ans int64) {
		for _, v := range arr {
			ans += int64(v)
		}
		return ans
	}
	n := len(nums)
	f := make([]int, n)
	copy(f, nums)
	ans := sum(f)
	for k := 1; k < n; k++ {
		for i := 0; i < n; i++ {
			f[i] = min(f[i], nums[(i+k)%n])
		}
		ans = min(ans, int64(k)*int64(x)+sum(f))
	}
	return ans
}

func main() {

}
