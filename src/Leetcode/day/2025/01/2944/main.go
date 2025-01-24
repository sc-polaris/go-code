package main

import (
	"math"
	"slices"
)

/*
	给你一个 下标从 1 开始的 整数数组 prices ，其中 prices[i] 表示你购买第 i 个水果需要花费的金币数目。

	水果超市有如下促销活动：
	· 如果你花费 prices[i] 购买了下标为 i 的水果，那么你可以免费获得下标范围在 [i + 1, i + i] 的水果。
	注意 ，即使你 可以 免费获得水果 j ，你仍然可以花费 prices[j] 个金币去购买它以获得它的奖励。

	请你返回获得所有水果所需要的 最少 金币数。
*/

/*
	我们需要解决的问题是：「获得第 1 个及其后面的水果所需要的最少金币数」。
	第 1 个水果一定要买，然后呢？
	第 2 个水果可以购买，也可以免费获得：
	· 如果购买，那么需要解决的问题为：「在购买第 2 个水果的前提下，获得第 2 个及其后面的水果所需要的最少金币数」。
	· 如果免费获得，那么根据题意，第 3 个水果必须购买，需要解决的问题为：「在购买第 3 个水果的前提下，获得第 3 个及其后面的水果所需要的最少金币数」。

	从上面的讨论可以知道，只需要一个 i 就能表达子问题，即定义 dfs(i) 表示在购买第 i 个水果的前提下，获得第 i 个及其
	后面的水果所需要的最少金币数。注意 i 从 1 开始。

	买第 i 个水果，那么从 i+1 到 2i 的水果都是免费的。枚举下一个购买的水果 j，问题变成：在购买第 j 个水果的前提下，获得第
	j 个及其后面的水果所需要的最少金币数，即 dfs(j)。

	j 的范围是 [i+1,2i+1]。其中 2i+1 表示免费获得从 i+1 到 2i 的所有水果，那么第 2i+1 个水果不能免费，一定要买。

*/

func minimumCoins(prices []int) int {
	n := len(prices)
	memo := make([]int, (n+1)/2)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i*2 >= n { // 此时后面的水果都可以免费获得了
			return prices[i-1] // i 从 1 开始
		}
		p := &memo[i]
		if *p != 0 {
			return *p
		}
		defer func() { *p = res }()
		res = math.MaxInt
		for j := i + 1; j <= i*2+1; j++ {
			res = min(res, dfs(j))
		}
		return res + prices[i-1]
	}
	return dfs(1)
}

/*
	递推
	f[i] 的定义和 dfs(i) 的定义是一样的，都表示在购买第 i 个水果的前提下，获得第 i 个及其后面的水果所需要的最少金币数。注意 i 从 1 开始。
							   2i+1
			f[i] = prices[i] + min f[j]
 							   j=i+1
	注：由于从比 i 更大的 j 转移过来，所以必须倒着计算 f。
	初始值：当 i≥⌊(n+1)/2⌋ 时，f[i]=prices[i]，翻译自递归边界 dfs(i)=prices[i]。
	答案：f[1]，翻译自递归入口 dfs(1)。
	注：由于传入的 prices 数组的下标是从 0 开始的，代码实现时下标要减一。
	代码实现时，可以直接把 prices 当作 f 数组。
*/

func minimumCoins2(prices []int) int {
	n := len(prices)
	for i := (n+1)/2 - 1; i > 0; i-- {
		prices[i-1] += slices.Min(prices[i : i*2+1])
	}
	return prices[0]
}

/*
	单调队列优化
	由于随着 i 的变小，j 的范围 [i+1,2i+1] 的左右边界也在变小，所以 [i+1,2i+1] 是一个向左的滑动窗口。
		2i+1
	计算 min f[j] 的过程本质上是在计算滑动窗口最小值
 		j=i+1
	下面代码中的队首在左边，队尾在右边。
*/

func minimumCoins3(prices []int) int {
	n := len(prices)
	type pair struct{ i, f int }
	q := []pair{{n + 1, 0}} // 哨兵
	for i := n; i > 0; i-- {
		for q[0].i > i*2+1 { // 右边离开窗口
			q = q[1:]
		}
		f := prices[i-1] + q[0].f
		for f <= q[len(q)-1].f {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, f}) // 左边进入窗口
	}
	return q[len(q)-1].f
}
