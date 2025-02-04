package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// 递归 O(N^2 logN)
func construct(grid [][]int) *Node {
	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, c0, c1 int) *Node {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				// 不是叶子节点
				if v != rows[0][c0] {
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &Node{
						Val:         true,
						IsLeaf:      false,
						TopLeft:     dfs(rows[:rMid], c0, cMid),
						TopRight:    dfs(rows[:rMid], cMid, c1),
						BottomLeft:  dfs(rows[rMid:], c0, cMid),
						BottomRight: dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		// 是叶子节点
		return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}

// 递归 + 二维前缀和优化 O(N^2)
func construct2(grid [][]int) *Node {
	n := len(grid)
	pre := make([][]int, n+1)
	pre[0] = make([]int, n+1)
	for i, row := range grid {
		pre[i+1] = make([]int, n+1)
		for j, v := range row {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}

	var dfs func(r0, c0, r1, c1 int) *Node
	dfs = func(r0, c0, r1, c1 int) *Node {
		total := pre[r1][c1] - pre[r1][c0] - pre[r0][c1] + pre[r0][c0]
		if total == 0 {
			return &Node{Val: false, IsLeaf: true}
		}
		if total == (r1-r0)*(c1-c0) {
			return &Node{Val: true, IsLeaf: true}
		}
		rMid, cMid := (r0+r1)/2, (c0+c1)/2
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     dfs(r0, c0, rMid, cMid),
			TopRight:    dfs(r0, cMid, rMid, c1),
			BottomLeft:  dfs(rMid, c0, r1, cMid),
			BottomRight: dfs(rMid, cMid, r1, c1),
		}
	}
	return dfs(0, 0, n, n)
}
