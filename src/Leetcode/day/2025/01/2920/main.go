package main

/*
	有一棵由 n 个节点组成的无向树，以 0  为根节点，节点编号从 0 到 n - 1 。给你一个长度为 n - 1 的二维 整数 数组 edges ，其中 edges[i] = [ai, bi]
	表示在树上的节点 ai 和 bi 之间存在一条边。另给你一个下标从 0 开始、长度为 n 的数组 coins 和一个整数 k ，其中 coins[i] 表示节点 i 处的金币数量。

	从根节点开始，你必须收集所有金币。要想收集节点上的金币，必须先收集该节点的祖先节点上的金币。

	节点 i 上的金币可以用下述方法之一进行收集：
	· 收集所有金币，得到共计 coins[i] - k 点积分。如果 coins[i] - k 是负数，你将会失去 abs(coins[i] - k) 点积分。
	· 收集所有金币，得到共计 floor(coins[i] / 2) 点积分。如果采用这种方法，节点 i 子树中所有节点 j 的金币数 coins[j] 将会减少至 floor(coins[j] / 2) 。
	返回收集 所有 树节点的金币之后可以获得的最大积分。
*/

/*
	记忆化搜索
	floor(coins[i] / 2) 等价于 coins[i] >> 1。
	右移运算是可以叠加的，即 (x >> 1) >> 1 等于 x >> 2。
	我们可以在递归的过程中，额外记录从根节点递归到当前节点的过程中，一共执行了多少次右移，也就是子树中的每个节点值需要右移的次数。

	故定义 dfs(i,j) 表示递归到以 i 为根的子树，在上面已经执行了 j 次右移的前提下，我们在这棵子树中最多可以得到多少积分。
	用「选或不选」来思考，即是否执行右移：
	· 不右移：答案为 (coins[i] >> j)−k 加上 i 的每个子树 ch 的 dfs(ch,j)。
	· 右移：答案为 coins[i] >> (j+1) 加上 i 的每个子树 ch 的 dfs(ch,j+1)。

	递归入口：dfs(0,0)。其中 i=0 表示根节点。一开始没有执行右移，所以 j=0。

	细节
	一个数最多右移多少次，就变成 0 了？
	设 w 是 coins[i] 的二进制长度，那么 coins[i] 右移 w 次后就是 0 了。
	在本题的数据范围下，w≤14。
	所以如果在递归过程中发现 j+1=14，就不执行右移，因为此时 dfs(ch,j+1) 子树中的每个节点值都要右移 14 次，算出的结果一定是 0。既然都知道递归的结果了，那就不需要递归了。
	此外，为避免错把父亲当作儿子，可以额外传入 fa 表示父节点，遍历 i 的邻居时，跳过邻居节点是 fa 的情况。
*/

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([][14]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, fa int) (res int) {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res1 := int(coins[i]>>j - k)
		res2 := int(coins[i] >> (j + 1))
		for _, ch := range g[i] {
			if ch != fa {
				res1 += dfs(ch, j, i) // 不右移
				if j < 13 {           // j+1 >= 14 相当于 res2 += 0 无需递归
					res2 += dfs(ch, j+1, i) // 右移
				}
			}
		}
		res = max(res1, res2)
		return
	}
	return dfs(0, 0, -1)
}

/*
	递推
	类似把记忆化搜索 1:1 翻译成递推的过程，我们也可以从下往上算。

	去掉参数 j，改成每个节点 i 返回一个长为 14 的列表 fi，其中 fi[j] 对应上面 dfs(i,j) 的计算结果。
	递推式为
				fi[j]=max (coins[i] >> j)−k+∑ch fch[j]
						  (coins[i] >> (j+1))+∑ch fch[j+1]
	把 ∑ch fch[j]累加到 s[j] 中，上式为
				fi[j]=max (coins[i] >> j)−k+s[j]
						  (coins[i] >> (j+1))+s[j+1]
	特判 j=13 的情况，上式为
				fi[13]=(coins[i] >> 13)−k+s[13]
*/

func maximumPoints2(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) [14]int
	dfs = func(x, fa int) (s [14]int) {
		for _, y := range g[x] {
			if y != fa {
				fy := dfs(y, x)
				for j, v := range fy {
					s[j] += v
				}
			}
		}
		for j := range 13 {
			s[j] = max((coins[x]>>j)-k+s[j], (coins[x]>>(j+1))+s[j+1])
		}
		s[13] += (coins[x] >> 13) - k
		return
	}
	return dfs(0, -1)[0]
}
