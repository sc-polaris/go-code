package main

import "math"

/*
	给你两个数组 nums 和 andValues，长度分别为 n 和 m。

	数组的 值 等于该数组的 最后一个 元素。

	你需要将 nums 划分为 m 个 不相交的连续子数组，对于第 ith 个子数组 [li, ri]，子数组元素的按位 AND 运算结果等于 andValues[i]，换句话说，
	所有的 1 <= i <= m，nums[li] & nums[li + 1] & ... & nums[ri] == andValues[i] ，其中 & 表示按位 AND 运算符。

	返回将 nums 划分为 m 个子数组所能得到的可能的 最小 子数组 值 之和。如果无法完成这样的划分，则返回 -1 。
*/

/*
	寻找子问题
	看示例 1，nums=[1,4,3,3,2], andValues=[0,3,3,2]。
	我们要解决的问题是，把 nums 划分成 4 个子数组所能得到的最小子数组值之和，其中每个子数组的 AND 值与 andValues 中的值一一对应。
	从 nums[0] 开始。考虑是否要把 nums[0] 作为子数组的最后一个数，分类讨论：
	· 不把 nums[1] 作为子数组的最后一个数，也就是 nums[1] 和后续元素在同一个子数组中，那么接下来需要解决的问题为：把 [3,3,2] 划分成 4 个子数组，且第一个子数组与
	  nums[0]&nums[1] 计算 AND 的值恰好等于 andValues[0]=0，其余子数组的 AND 值分别等于 3,3,2，在满足该条件的情况下，所能得到的最小子数组值之和。注意剩余元素
	  只有 3 个，没法分成 4 个子数组。
	· 把 nums[1] 作为子数组的最后一个数，注意我们并不需要知道这个子数组的前面具体有哪些数，只需要知道前面的元素的 AND 值等于 1。由于nums[0]&nums[1]=1&4=0=andValues[0]，
	  符合题目要求，可以划分。接下来需要解决的问题为：把 [3,3,2] 划分成 3 个子数组，子数组的 AND 值分别等于 3,3,2，在满足该条件的情况下，所能得到的最小子数组值之和。

	状态定义与状态转移方程
	递归需要哪些参数？
	1. 需要知道当前考虑到 nums 的哪个数，其下标记作 i。
	2. 需要知道当前划分的子数组对应着 andValues 的哪个数，其下标记作 j。也可以理解为前面已经划分了 j 段。
	3. 需要知道当前划分的子数组，在 i 左边的那些元素的 AND 值，记作 and。再次强调，我们并不需要知道 i 左边具体有哪些数，只需要知道左边那些数的 AND 值是多少即可。

	于是，定义 dfs(i,j,and) 表示从左往右划分，目前考虑到 nums[i]，已经划分了 j 段，且当前待划分的这一段已经参与 AND 运算的结果为 and，在这种情况下，剩余元素划分得到的最小和。

	首先把 and 与 nums[i] 计算 AND。
	用「选或不选」的思想分类讨论：
	· 不划分：继续向右递归 dfs(i+1,j,and)。
	· 划分：如果 and=andValues[j]，那么可以划分，即 dfs(i+1,j+1,−1)+nums[i]。这里令 and=−1 是因为 −1 的二进制全为 1，与任何数 x 的 AND 都是 x，适合用来计算新子数组的 AND 值。
	· 这两种情况取最小值，就得到了 dfs(i,j,and)。

	递归边界：
	· 如果 n−i<m−j，那么剩余元素不足，无法划分，返回 ∞。
	· 如果 j=m 且 i<n，还有元素没有划分，返回 ∞。
	· 如果 j=m 且 i=n，划分成功，返回 0。

	递归入口：dfs(0,0,−1)，即答案。如果答案是 ∞ 则返回 −1。
*/

func minimumValueSum(nums []int, andValues []int) int {
	const inf = math.MaxInt / 2 // 除 2 防止下面 +nums[i] 溢出
	n, m := len(nums), len(andValues)
	type args struct{ i, j, and int }
	memo := make(map[args]int)
	var dfs func(int, int, int) int
	dfs = func(i int, j int, and int) int {
		if n-i < m-j { // 剩余元素不足
			return inf
		}
		if j == m { // 分了 m 段
			if i == n {
				return 0
			}
			return inf
		}

		and &= nums[i]
		p := args{i, j, and}
		if res, ok := memo[p]; ok { // 之前计算过
			return res
		}
		res := dfs(i+1, j, and)  // 不划分
		if and == andValues[j] { // 划分，nums[i] 是这一段的最后一个数
			res = min(res, dfs(i+1, j+1, -1)+nums[i])
		}
		memo[p] = res
		return res
	}
	ans := dfs(0, 0, -1)
	if ans == inf {
		return -1
	}
	return ans
}
