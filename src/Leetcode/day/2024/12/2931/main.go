package main

import "slices"

/*
	给你一个下标从 0 开始大小为 m * n 的整数矩阵 values ，表示 m 个不同商店里 m * n 件不同的物品。每个商店有 n 件物品，第 i 个商店的第 j 件物品的价值为 values[i][j] 。
	除此以外，第 i 个商店的物品已经按照价值非递增排好序了，也就是说对于所有 0 <= j < n - 1 都有 values[i][j] >= values[i][j + 1] 。
	每一天，你可以在一个商店里购买一件物品。具体来说，在第 d 天，你可以：
	· 选择商店 i 。
	· 购买数组中最右边的物品 j ，开销为 values[i][j] * d 。换句话说，选择该商店中还没购买过的物品中最大的下标 j ，并且花费 values[i][j] * d 去购买。
	注意，所有物品都视为不同的物品。比方说如果你已经从商店 1 购买了物品 0 ，你还可以在别的商店里购买其他商店的物品 0 。

	请你返回购买所有 m * n 件物品需要的 最大开销 。
*/

/*
	根据 排序不等式，values[i][j] 越小的数，应该越早购买。
*/

// 排序
func maxSpending(values [][]int) (ans int64) {
	m, n := len(values), len(values[0])
	a := make([]int, 0, m*n)
	for _, row := range values {
		a = append(a, row...)
	}
	slices.Sort(a)
	for i, x := range a {
		ans += int64(x) * int64(i+1)
	}
	return
}
