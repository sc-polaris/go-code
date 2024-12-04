package main

/*
	有一个 8 x 8 的棋盘，它包含 n 个棋子（棋子包括车，后和象三种）。给你一个长度为 n 的字符串数组 pieces ，其中 pieces[i] 表示第 i 个棋子的类型（车，后或象）。除此以外，还给你一个长度为 n
	的二维整数数组 positions ，其中 positions[i] = [ri, ci] 表示第 i 个棋子现在在棋盘上的位置为 (ri, ci) ，棋盘下标从 1 开始。

	棋盘上每个棋子都可以移动 至多一次 。每个棋子的移动中，首先选择移动的 方向 ，然后选择 移动的步数 ，同时你要确保移动过程中棋子不能移到棋盘以外的地方。棋子需按照以下规则移动：
	· 车可以 水平或者竖直 从 (r, c) 沿着方向 (r+1, c)，(r-1, c)，(r, c+1) 或者 (r, c-1) 移动。
	· 后可以 水平竖直或者斜对角 从 (r, c) 沿着方向 (r+1, c)，(r-1, c)，(r, c+1)，(r, c-1)，(r+1, c+1)，(r+1, c-1)，(r-1, c+1)，(r-1, c-1) 移动。
	· 象可以 斜对角 从 (r, c) 沿着方向 (r+1, c+1)，(r+1, c-1)，(r-1, c+1)，(r-1, c-1) 移动。
	移动组合 包含所有棋子的 移动 。每一秒，每个棋子都沿着它们选择的方向往前移动 一步 ，直到它们到达目标位置。所有棋子从时刻 0 开始移动。如果在某个时刻，两个或者更多棋子占据了同一个格子，那么这个移动组合 不有效 。
	请你返回 有效 移动组合的数目。
	注意：
	· 初始时，不会有两个棋子 在 同一个位置 。
	· 有可能在一个移动组合中，有棋子不移动。
	· 如果两个棋子 直接相邻 且两个棋子下一秒要互相占据对方的位置，可以将它们在同一秒内 交换位置 。
*/

/*
	核心思路
	1. 预处理每个棋子的所有合法移动。
	2. 写一个回溯，暴力枚举每个棋子的每个合法移动，如果这些棋子没有重叠在一起，则答案加一。

	细节
	具体来说，合法移动包含：
	· 棋子的初始位置 (x0,y0)。
	· 棋子的移动方向 (dx,dy)。
	· 棋子的移动次数 step。
	在回溯时，可以剪枝：如果当前棋子的当前这个合法移动，与前面的棋子冲突，即同一时刻两个棋子重叠，那么不往下递归，枚举当前棋子的下一个合法移动。
*/

type move struct {
	x0, y0 int // 起点
	dx, dy int // 移动方向
	step   int // 移动步数
}

type dir struct{ x, y int }

var dirs = []dir{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}} // 上下左右 + 斜向
var pieceDirs = map[byte][]dir{'r': dirs[:4], 'b': dirs[4:], 'q': dirs}                // 车、象、后

// 计算位于 (x0,y0) 的棋子在 dirs 这些方向上的所有合法移动
func generateMoves(x0, y0 int, dirs []dir) []move {
	const size = 8
	moves := []move{{x0, y0, 0, 0, 0}} // 原地不动
	for _, d := range dirs {
		// 往 d 方向走 1,2,3,... 步
		x, y := x0+d.x, y0+d.y
		for step := 1; 0 < x && x <= size && 0 < y && y <= size; step++ {
			moves = append(moves, move{x0, y0, d.x, d.y, step})
			x += d.x
			y += d.y
		}
	}
	return moves
}

// 判断两个移动是否合法，即不存在同一时刻两个棋子重叠的情况
func isValid(m1, m2 move) bool {
	x1, y1 := m1.x0, m1.y0
	x2, y2 := m2.x0, m2.y0
	for i := range max(m1.step, m2.step) {
		// 每一秒走一步
		if i < m1.step {
			x1 += m1.dx
			y1 += m1.dy
		}
		if i < m2.step {
			x2 += m2.dx
			y2 += m2.dy
		}
		if x1 == x2 && y1 == y2 { // 重叠
			return false
		}
	}
	return true
}

func countCombinations(pieces []string, positions [][]int) (ans int) {
	n := len(pieces)
	// 预处理所有合法移动
	allMoves := make([][]move, n)
	for i, pos := range positions {
		allMoves[i] = generateMoves(pos[0], pos[1], pieceDirs[pieces[i][0]])
	}

	path := make([]move, n) // 注意 path 的长度是固定的
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}
	outer:
		// 枚举当前棋子的所有合法移动
		for _, move1 := range allMoves[i] {
			// 判断合法移动 move1 是否有效
			for _, move2 := range path[:i] {
				if !isValid(move1, move2) {
					continue outer // 无效，枚举下一个 move1
				}
			}
			path[i] = move1 // 直接覆盖，无需恢复现场
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}
