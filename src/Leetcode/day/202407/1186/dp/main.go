package dp

import "math"

func maximumSum(arr []int) int {
	ans := math.MinInt
	f := make([][2]int, len(arr)+1)
	f[0] = [2]int{math.MinInt / 2, math.MinInt / 2}
	for i, x := range arr {
		f[i+1][0] = max(f[i][0], 0) + x
		f[i+1][1] = max(f[i][1]+x, f[i][0])
		ans = max(ans, f[i+1][0], f[i+1][1])
	}
	return ans
}

/*
	计算顺序！必须先算 f[1] 再算 f[0]。如果先算 f[0] 再算 f[1]，那么在计算 f[1] 时，相当于用到的不是原来的 f[i][0]，而是新算出来的 f[i+1][0]。
*/

func maximumSum2(arr []int) int {
	ans := math.MinInt / 2
	f0, f1 := ans, ans
	for _, x := range arr {
		f1 = max(f1+x, f0)
		f0 = max(f0, 0) + x
		ans = max(ans, f0, f1)
	}
	return ans
}
