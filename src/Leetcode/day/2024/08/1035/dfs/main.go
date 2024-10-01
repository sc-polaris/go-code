package dfs

/*
	在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

	现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：
	· nums1[i] == nums2[j]
	· 且绘制的直线不与任何其他连线（非水平线）相交。
	请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

	以这种方法绘制线条，并返回可以绘制的最大连线数。
*/

/*
	s = [1,4,2] t = [1,2,4]
	用「选或不选」分类讨论：
	· 不选 s[2]=2，那么需要解决的问题为：s=[1,4],t=[1,2,4] 的最大连接数
	· 不选 t[2]=4，那么需要解决的问题为：s=[1,4,2],t=[1,2] 的最大连接数
	除此以外，对于 s = [1,4]，t=[1,2,4]，由于 s[1]=t[2]=4，我们都可以选，可就是在两个数字之间连线，问题变成：s=[1],t=[1,2] 的最大连线数。

	定义 dfs(i,j) 表示 s[0] 到 s[i] 与 t[0] 到 t[j] 之间的最大连线数。
	状态转移方程
						dfs(i-1,j-1) + 1			s[i] == t[j]
			dfs(i,j) =
						max(dfs(i-1,j),dfs(i,j-1))	s[i] != t[j]

	递归边界：dfs(−1,j)=dfs(i,−1)=0。当其中一个数组为空时，连线数等于 0。
	递归入口：dfs(n−1,m−1)，也就是答案。
*/

func maxUncrossedLines(s []int, t []int) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if s[i] == t[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(n-1, m-1)
}
