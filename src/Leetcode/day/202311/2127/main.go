package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	indeg := make([]int, n)
	for _, f := range favorite {
		indeg[f]++ // 统计基环树每个节点的入度
	}

	depth := make([]int, n)
	var q []int
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 { // 拓扑排序
		x := q[0]
		q = q[1:]
		y := favorite[x] // x只有一条出边
		depth[y] = depth[x] + 1
		if indeg[y]--; indeg[y] == 0 {
			q = append(q, y)
		}
	}

	maxRingSize, sumChainSize := 0, 0
	for i, d := range indeg {
		if d == 0 {
			continue
		}

		// 遍历基环上的点
		indeg[i] = 0  // 将基环上的点的入度标记为0，避免重复访问
		ringSize := 1 // 基环长度
		for x := favorite[i]; x != i; x = favorite[x] {
			indeg[x] = 0 // 将基环上的点的入度标记为0，避免重复访问
			ringSize++
		}
		if ringSize == 2 {
			sumChainSize += depth[i] + depth[favorite[i]] + 2 // 累加两条最长链的长度
		} else {
			maxRingSize = max(maxRingSize, ringSize) // 去所有基环长度的最大值
		}
	}
	return max(maxRingSize, sumChainSize)
}

func main() {

}
