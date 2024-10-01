package main

import (
	"math/bits"
)

/*
	组合数学

	假设使用了 m 次操作一，j 次操作二，那么有
			1+2^0+2^1+2^2+⋯+2^(j−1)−m = k
	即
				m = 2^j - k
	注意上式当 j=0 时也是成立的。

	由于操作一不能连续使用，我们需要在这 j 次操作二前后，以及相邻两次操作二的空隙中，插入 m 个操作一，所以方案数等于从 j+1 个物品中选出 m 个物品的方案数。

	其中 0 ≤ m ≤ j+1。根据题目的数据范围，j 至多枚举到 29。

	8!		= 1x...x8  				= 8 7 6
	3!5!	= 1x2x3 x 1x2x3x4x5		  123
*/

const mx = 31

var c [mx][mx]int

func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

//func waysToReachStair(k int) (ans int) {
//	for j := 0; j < 30; j++ {
//		m := 1<<j - k
//		if 0 <= m && m <= j+1 {
//			ans += c[j+1][m]
//		}
//	}
//	return
//}

/*
	优化

	只需要从首个满足 2^j ≥ k 的 j 开始枚举，相当于 j 至少是 max(k−1,0) 的二进制长度（0 的二进制长度为 0）。
	当 2^j−k > j+1 时，停止循环。
*/

func waysToReachStair(k int) (ans int) {
	for j := bits.Len(uint(max(k-1, 0))); 1<<j-k <= j+1; j++ {
		ans += c[j+1][1<<j-k]
	}
	return
}
