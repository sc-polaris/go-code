package main

import (
	"math"
	"slices"
)

/*
	有一根长度为 n 个单位的木棍，棍上从 0 到 n 标记了若干位置。例如，长度为 6 的棍子可以标记如下：
	给你一个整数数组 cuts ，其中 cuts[i] 表示你需要将棍子切开的位置。

	你可以按顺序完成切割，也可以根据需要更改切割的顺序。

	每次切割的成本都是当前要切割的棍子的长度，切棍子的总成本是历次切割成本的总和。对棍子进行切割将会把
	一根木棍分成两根较小的木棍（这两根木棍的长度和就是切割前木棍的长度）。请参阅第一个示例以获得更直观的解释。

	返回切棍子的 最小总成本 。
*/

/*
	一、寻找子问题
	示例 1 的 cuts=[1,3,4,5]，为方便描述，把 0 和 n=7 也视作切开的位置（木棍端点），得到 cuts=[0,1,3,4,5,7]。
	我们要解决的问题（原问题）是：
		· 切割一根左端点为 cuts[0]=0，右端点为 cuts[5]=7 的棍子的最小成本。
	第一刀切在哪？枚举：
		· 在 cuts[1]=1 切一刀，木棍分成两段。第一段左端点为 cuts[0]=0，右端点为 cuts[1]=1；第二段左端点为 cuts[1]=1，右端点为 cuts[5]=7。
		· 在 cuts[2]=3 切一刀，木棍分成两段。第一段左端点为 cuts[0]=0，右端点为 cuts[2]=3；第二段左端点为 cuts[2]=3，右端点为 cuts[5]=7。
		· 在 cuts[3]=4 切一刀，木棍分成两段。第一段左端点为 cuts[0]=0，右端点为 cuts[3]=4；第二段左端点为 cuts[3]=4，右端点为 cuts[5]=7。
		· 在 cuts[4]=5 切一刀，木棍分成两段。第一段左端点为 cuts[0]=0，右端点为 cuts[4]=5；第二段左端点为 cuts[4]=5，右端点为 cuts[5]=7。
	接下来，继续计算这两段木棍各自的最小切割成本。同样地，枚举切割的位置。依此类推。
	这些问题都是和原问题相似的、规模更小的子问题，可以用递归解决。

	二、状态定义与状态转移方程
	根据上面的讨论，我们需要在递归过程中，知道当前切的这根棍子，左端点在哪，右端点在哪。
	因此，定义状态为 dfs(i,j)，表示切割一根左端点为 cuts[i]，右端点为 cuts[j] 的棍子的最小成本。
	枚举在 cuts[k] 处切一刀，其中 k=i+1,i+2,…,j−1，木棍变成两段：
		· 第一段左端点为 cuts[i]，右端点为 cuts[k]，切割这段木棍的最小成本为 dfs(i,k)。
		· 第二段左端点为 cuts[k]，右端点为 cuts[j]，切割这段木棍的最小成本为 dfs(k,j)。
		· 成本之和为 dfs(i,k)+dfs(k,j)，再算上切割之前木棍的长度 cuts[j]−cuts[i]，得到
								dfs(i,k)+dfs(k,j)+cuts[j]−cuts[i]
	枚举 k=i+1,i+2,…,j−1，所有成本取最小值，得
				  j-1
		dfs(i,j)= min dfs(i,k)+dfs(k,j)+cuts[j]−cuts[i]
				  k=i+1
	其中 cuts[j]−cuts[i] 与 k 无关，可以提到循环外面。
	递归边界：dfs(i,i+1)=0。此时木棍中没有要切割的位置，所以切割成本为 0。
	递归入口：dfs(0,m−1)，也就是答案。其中 m 是添加了 0 和 n 之后的 cuts 数组的长度。

*/

func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	slices.Sort(cuts)

	m := len(cuts)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i+1 == j { // 无需切割
			return 0
		}
		p := &memo[i][j]
		if *p != 0 {
			return *p
		}
		res := math.MaxInt
		for k := i + 1; k < j; k++ {
			res = min(res, dfs(i, k)+dfs(k, j))
		}
		*p = res + cuts[j] - cuts[i] // 记忆化
		return *p
	}
	return dfs(0, m-1)
}

func minCost2(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	slices.Sort(cuts)

	m := len(cuts)
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, m)
	}
	for i := m - 3; i >= 0; i-- {
		for j := i + 2; j < m; j++ {
			res := math.MaxInt
			for k := i + 1; k < j; k++ {
				res = min(res, f[i][k]+f[k][j])
			}
			f[i][j] = res + cuts[j] - cuts[i]
		}
	}
	return f[0][m-1]
}
