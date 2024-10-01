package main

import "slices"

/*
	一个大小为 k 的连通块内，如果只有一个节点 x 被感染（x 在 initial 中），那么移除 x 后，这个连通块不会被感染，
	从而让 M(initial) 减少 k。

	而如果连通块内至少有两个节点被感染，无论移除哪个节点，仍然会导致连通块被感染，M(initial) 不变。

	所以我们要找的是只包含一个被感染节点的连通块，并且这个连通块越大越好。

	算法如下：
	1. 遍历 initial 中的节点 x。
	2. 如果 x 没有被访问过，那么从 x 开始 DFS，同时用一个 vis 数组标记访问过的节点。
	3. DFS 过程中，统计连通块的大小 size。
	4. DFS 过程中，记录访问到的在 initial 中的节点。
	5. DFS 结束后，如果发现该连通块只有一个在 initial 中的节点，并且该连通块的大小比最大的连通块更大，那么更新最大连通块的大小，
	   已经答案节点 x。如果一样大，就更新答案节点的最小值。
	6. 最后，如果没有找到符合要求的节点，返回 min(initial)；否则返回答案节点。

	如何表达出 「连通块内有一个或多个被感染的节点」呢？要记录被感染的节点列表吗？
	其实无需记录节点列表，而是用如下状态机：

			-1   -->     x       -->  -2
     	  初始状态      找到一个        找到多个
	· 初始状态为 -1。
	· 如果状态是 -1，在找到被感染的节点 x 后，状态变为 x。
	· 如果状态是非负数 x，在找到另一个被感染的节点后，状态变为 -2。如果状态已经是 -2，则不变。

	此外，可以用一个哈希表或者布尔数组，记录哪些点在 initial 中，从而在 DFS 中快速判断当前节点是否在 initial 中。
*/

func minMalwareSpread(graph [][]int, initial []int) int {
	vis := make([]bool, len(graph))
	isInitial := make([]bool, len(graph))
	for _, x := range initial {
		isInitial[x] = true
	}

	var nodeId, size int
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		// 按照状态机更新 nodeId
		if nodeId != -2 && isInitial[x] {
			if nodeId < 0 {
				nodeId = x
			} else {
				nodeId = -2
			}
		}
		for y, conn := range graph[x] {
			if conn == 1 && !vis[y] {
				dfs(y)
			}
		}
	}

	ans := -1
	maxSize := 0
	for _, x := range initial {
		if vis[x] {
			continue
		}
		nodeId = -1
		size = 0
		dfs(x)
		if nodeId >= 0 && (size > maxSize || size == maxSize && nodeId < ans) {
			ans = nodeId
			maxSize = size
		}
	}
	if ans < 0 {
		return slices.Min(initial)
	}
	return ans
}
