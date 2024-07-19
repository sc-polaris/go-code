package main

/*
	给你一个长度为 n 的二进制数组 possible 。
	Alice 和 Bob 正在玩一个有 n 个关卡的游戏，游戏中有一些关卡是 困难 模式，其他的关卡是 简单 模式。如果 possible[i] == 0 ，那么第 i 个关卡
	是 困难 模式。一个玩家通过一个简单模式的关卡可以获得 1 分，通过困难模式的关卡将失去 1 分。

	游戏的一开始，Alice 将从第 0 级开始 按顺序 完成一些关卡，然后 Bob 会完成剩下的所有关卡。
	假设两名玩家都采取最优策略，目的是 最大化 自己的得分，Alice 想知道自己 最少 需要完成多少个关卡，才能获得比 Bob 更多的分数。
	请你返回 Alice 获得比 Bob 更多的分数所需要完成的 最少 关卡数目，如果 无法 达成，那么返回 -1 。
	注意，每个玩家都至少需要完成 1 个关卡。
*/

/*
	题意：把 0 当作 -1，找一个最短前缀，其元素和大于剩余元素和。
	设 possible 的元素和为 s（把 0 视作 -1）。
	枚举 x = possible[i]，同时计算前缀和 pre，那么剩余元素为 s - pre
	如果
					pre > s - pre
	即
					pre * 2 > s
	就返回 i + 1，即前缀长度。

	代码实现时，计算 pre 可以把 1 视作 2，把 0 视作 -2，这样无需计算乘 2。
*/

func minimumLevels(possible []int) int {
	n := len(possible)
	// s = cnt1 - cnt0 = cnt1 - (n - cnt1) = = cnt1 * 2 - n
	s := 0
	for _, x := range possible {
		s += x
	}
	s = s*2 - n
	pre := 0
	for i, x := range possible[:n-1] {
		pre += x*4 - 2
		if pre > s {
			return i + 1
		}
	}
	return -1
}
