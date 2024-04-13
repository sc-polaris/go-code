package main

/*
	标记所有 edges[i][1]，这些对五都不是冠军。

	然后遍历每个节点 i，如果恰好有一个 i 没有被标记，说明没有可以击败 i 的队伍，i 队是冠军。否则返回 -1。
*/

func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true // 不是冠军
	}

	ans := -1
	for i, w := range weak {
		if w {
			continue
		}
		if ans != -1 {
			return -1 // 冠军只能有一个
		}
		ans = i
	}
	return ans
}
