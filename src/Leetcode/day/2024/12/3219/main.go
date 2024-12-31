package main

import "slices"

/*
	有一个 m x n 大小的矩形蛋糕，需要切成 1 x 1 的小块。

	给你整数 m ，n 和两个数组：
	· horizontalCut 的大小为 m - 1 ，其中 horizontalCut[i] 表示沿着水平线 i 切蛋糕的开销。
	· verticalCut 的大小为 n - 1 ，其中 verticalCut[j] 表示沿着垂直线 j 切蛋糕的开销。

	一次操作中，你可以选择任意不是 1 x 1 大小的矩形蛋糕并执行以下操作之一：
	1. 沿着水平线 i 切开蛋糕，开销为 horizontalCut[i] 。
	2. 沿着垂直线 j 切开蛋糕，开销为 verticalCut[j] 。

	每次操作后，这块蛋糕都被切成两个独立的小蛋糕。
	每次操作的开销都为最开始对应切割线的开销，并且不会改变。
	请你返回将蛋糕全部切成 1 x 1 的蛋糕块的 最小 总开销。
*/

/*
	逆向思维 最小生成树 kruskal 算法
	根据最小生成树的 Kruskal 算法，先把边权从小到大排序，然后遍历边，如果边的两个点属于不同连通块，则合并。

	我们用双指针计算答案：
	1. 从小到大排序两个数组。初始化 i=j=0。
	2. 如果 horizontalCut[i]<verticalCut[j]，把 n−j 条边权为 horizontalCut[i] 的边加入答案，然后 i 加一。
	3. 否则，把 m−i 条边权为 verticalCut[j] 的边加入答案，然后 j 加一。
	4. 循环次数为两个数组的长度之和，即 (m−1)+(n−1)=m+n−2。
*/

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) (ans int64) {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	i, j := 0, 0
	for range m + n - 2 {
		if j == n-1 || i < m-1 && horizontalCut[i] < verticalCut[j] {
			ans += int64(horizontalCut[i] * (n - j))
			i++
		} else {
			ans += int64(verticalCut[j] * (m - i))
			j++
		}
	}
	return
}
