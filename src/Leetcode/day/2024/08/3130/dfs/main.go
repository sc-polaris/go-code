package dfs

/*
	给你 3 个正整数 zero ，one 和 limit 。
	一个二进制数组 arr 如果满足以下条件，那么我们称它是 稳定的 ：
	· 0 在 arr 中出现次数 恰好 为 zero 。
	· 1 在 arr 中出现次数 恰好 为 one 。
	· arr 中每个长度超过 limit 的子数组 都 同时 包含 0 和 1 。
	请你返回 稳定 二进制数组的 总 数目。
	由于答案可能很大，将它对 10^9 + 7 取余 后返回。
*/

/*
	先解释 limit，意思是数组中至多有连续 limit 个 0，且至多有连续 limit 个 1。
	看上去，定义	dfs(i,j)	表示用 i 个 0 和 j 个 1 构造稳定数组的方案数？但这样定义不方便计算 limit 带来的影响。
	改成定义		dfs(i,j,k)	表示用 i 个 0 和 j 个 1 构造稳定数组的方案数，其中第 i+j 个位置要填 k，其中 k 为 0 或 1。
	考虑 dfs(i,j,0) 怎么算。现在，第 i+j 个位置填的是 0，考虑第 i+j−1 个位置要填什么：
	· 填 0，问题变成 dfs(i-1,j,0)。
	· 填 1，问题变成 dfs(i-1,j,1)。
	看上去，把这两种情况加起来，我们就得到了 dfs(i,j,0)。
	但是，dfs(i−1,j,0) 包含了「最后连续 limit 个位置填 0」的方案，如果在这个方案末尾再加一个 0，就有连续 limit+1 个 0 了，这是不合法的，要减掉。
	dfs(i−1,j,0) 中的「最后连续 limit 个位置填 0」的方案有多少个？

	因为 dfs 的定义是稳定数组的方案数，只包含合法方案，所以在最后连续 limit 个位置填 0 的情况下，倒数第 limit+1 个位置一定要填 1，这有 dfs(i−limit−1,j,1) 种方案。
	对于 dfs(i,j,0) 来说，这 dfs(i−limit−1,j,1) 个方案就是不合法方案了，要减掉，得
					dfs(i,j,0)=dfs(i−1,j,0)+dfs(i−1,j,1)−dfs(i−limit−1,j,1)
	同理得
					dfs(i,j,1)=dfs(i,j−1,0)+dfs(i,j−1,1)−dfs(i,j−limit−1,0)

	递归边界 1：如果 i<0 或者 j<0，返回 0。也可以在递归 dfs(i−limit−1,j,1) 前判断 i>limit，在递归 dfs(i,j−limit−1,0) 前判断 j>limit。下面代码在递归前判断。
	递归边界 2：如果 i=0，那么当 k=1 且 j≤limit 的情况下返回 1，否则返回 0；如果 j=0，那么当 k=0 且 i≤limit 的情况下返回 1，否则返回 0。
	递归入口  ：dfs(zero,one,0)+dfs(zero,one,1)，即答案。
*/

func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i == 0 { // 递归边界
			if k == 1 && j <= limit {
				return 1
			}
			return
		}
		if j == 0 { // 递归边界
			if k == 0 && i <= limit {
				return 1
			}
			return
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		if k == 0 {
			// +mod 保证答案非负
			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
			if i > limit {
				res = (res - dfs(i-limit-1, j, 1) + mod) % mod
			}
		} else {
			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1) + mod) % mod
			if j > limit {
				res = (res - dfs(i, j-limit-1, 0) + mod) % mod
			}
		}
		*p = res
		return
	}
	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
}
