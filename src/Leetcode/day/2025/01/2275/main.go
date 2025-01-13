package main

import (
	"math/bits"
	"slices"
)

/*
	对数组 nums 执行 按位与 相当于对数组 nums 中的所有整数执行 按位与 。
	· 例如，对 nums = [1, 5, 3] 来说，按位与等于 1 & 5 & 3 = 1 。
	· 同样，对 nums = [7] 而言，按位与等于 7 。
	给你一个正整数数组 candidates 。计算 candidates 中的数字每种组合下 按位与 的结果。

	返回按位与结果大于 0 的 最长 组合的长度。
*/

/*
	枚举比特位

	既然要求 AND 大于 0，那么这个 AND 值中，一定有一个比特位是 1。
	枚举这个比特位：
	· 如果 AND 的最低位是 1，最多可以从 candidates 中选多少个数？
	· 如果 AND 的次低位是 1，最多可以从 candidates 中选多少个数？
	· 依此类推

	如果最低位是 1，那么从 candidates 中选的数，最低位也必须是 1。这样问题就变成：
	· candidates 中有多少个数，最低位是 1？
	遍历 candidates 即可算出。
*/

func largestCombination(candidates []int) (ans int) {
	m := bits.Len(uint(slices.Max(candidates)))
	for i := range m {
		cnt := 0
		for _, x := range candidates {
			cnt += x >> i & 1
		}
		ans = max(ans, cnt)
	}
	return
}

/*
	一次遍历
	创建一个大小为 24（10^7 的二进制长度）的 cnt 数组，统计每个比特位上的 1 的个数。
*/

func largestCombination2(candidates []int) int {
	cnt := [24]int{}
	for _, x := range candidates {
		for i := 0; x > 0; i++ {
			cnt[i] += x & 1
			x >>= 1
		}
	}
	return slices.Max(cnt[:])
}
