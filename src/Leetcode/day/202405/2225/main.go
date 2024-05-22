package main

import "slices"

/*
	给你一个整数数组 matches 其中 matches[i] = [winneri, loseri] 表示在一场比赛中 winneri 击败了 loseri 。

	返回一个长度为 2 的列表 answer ：
	· answer[0] 是所有 没有 输掉任何比赛的玩家列表。
	· answer[1] 是所有恰好输掉 一场 比赛的玩家列表。
	两个列表中的值都应该按 递增 顺序返回。

	注意：
	· 只考虑那些参与 至少一场 比赛的玩家。
	· 生成的测试用例保证 不存在 两场比赛结果 相同 。
*/

func findWinners(matches [][]int) [][]int {
	lossCount := make(map[int]int)
	for _, m := range matches {
		if lossCount[m[0]] == 0 {
			lossCount[m[0]] = 0
		}
		lossCount[m[1]]++
	}

	ans := make([][]int, 2)
	for player, cnt := range lossCount {
		if cnt < 2 {
			ans[cnt] = append(ans[cnt], player)
		}
	}

	slices.Sort(ans[0])
	slices.Sort(ans[1])
	return ans
}
