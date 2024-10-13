package main

import "math"

/*
	给你 2 枚相同 的鸡蛋，和一栋从第 1 层到第 n 层共有 n 层楼的建筑。

	已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都 会碎 ，从 f 楼层或比它低 的楼层落下的鸡蛋都 不会碎 。

	每次操作，你可以取一枚 没有碎 的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。

	请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
*/

/*
	如果鸡蛋碎了，那么接下来只能依次在 1,2,3,j−1 楼扔第二枚鸡蛋，最坏情况下，总共要操作 1+(j−1)=j 次。
	如果鸡蛋没碎，那么接下来可以在 j+1 到 i 楼中继续扔第一枚鸡蛋，这等价于在一栋有 i−j 层楼的建筑中扔鸡蛋的子问题，即 dfs(i−j)，将其加一即为总操作次数。
*/

func twoEggDrop(n int) int {
	var memo [1001]int
	var dfs func(int) int
	dfs = func(n int) int {
		if n == 0 {
			return 0
		}
		p := &memo[n]
		if *p > 0 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		for j := 1; j <= n; j++ {
			res = min(res, max(j, dfs(n-j)+1))
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n)
}

func twoEggDrop2(n int) int {
	var f [1001]int
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt
		for j := 1; j <= i; j++ {
			f[i] = min(f[i], max(j, f[i-j]+1))
		}
	}
	return f[n]
}
