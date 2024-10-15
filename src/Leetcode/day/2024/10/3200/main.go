package main

import "math"

/*
	给你两个整数 red 和 blue，分别表示红色球和蓝色球的数量。你需要使用这些球来组成一个三角形，
	满足第 1 行有 1 个球，第 2 行有 2 个球，第 3 行有 3 个球，依此类推。

	每一行的球必须是 相同 颜色，且相邻行的颜色必须 不同。

	返回可以实现的三角形的 最大 高度。
*/

// 枚举
func maxHeightOfTriangle(red int, blue int) int {
	cnt := [2]int{}
	for i := 1; ; i++ {
		cnt[i%2] += i
		// 这两个条件都成立，说明无法把第 i 行填满，返回 i−1。
		if (cnt[0] > red || cnt[1] > blue) && (cnt[0] > blue || cnt[1] > red) {
			return i - 1
		}
	}
}

/*
    数学公式
	核心思路
	奇数行放红球，偶数行放蓝球；或者奇数行放蓝球，偶数行放红球。

	计算最多能放多少排。两种情况取最大值。

	奇数行
	设奇数行有 k 行，那么需要
			1+3+5+⋯+(2k−1)=k^2
	个球。（等差数列求和公式n(a1+an)/2）
	假设我们有 n 个球，那么有
				n >= k^2
	解得
				k <= ⌊sqrt(n)⌋
	偶数行
	设偶数行有 k 行，那么需要
			2+4+6+⋯+2k=k^2+k
	个球。（等差数列求和公式）
	假设我们有 n 个球，那么有
				n >= k^2+k
	解得
				k <= ⌊(sqrt(4n+1)-1)/2⌋
	答案
	设有 odd 个奇数行，even 个偶数行，那么总行数为
				2even+1,	odd > even
				2odd,		otherwise

*/

func maxHeightOfTriangle2(red int, blue int) int {
	f := func(n, m int) int {
		odd := int(math.Sqrt(float64(n)))
		even := int((math.Sqrt(float64(4*m+1)) - 1) / 2)
		if odd > even {
			return 2*even + 1
		}
		return 2 * odd
	}
	return max(f(red, blue), f(blue, red))
}
