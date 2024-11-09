package main

/*
	给你一个 n x n 的二维数组 grid，它包含范围 [0, n2 - 1] 内的不重复元素。

	实现 neighborSum 类：
	· neighborSum(int [][]grid) 初始化对象。
	· int adjacentSum(int value) 返回在 grid 中与 value 相邻的元素之和，相邻指的是与 value 在上、左、右或下的元素。
	· int diagonalSum(int value) 返回在 grid 中与 value 对角线相邻的元素之和，对角线相邻指的是与 value 在左上、右上、左下或右下的元素。
*/

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}

type NeighborSum [][2]int

func Constructor(grid [][]int) NeighborSum {
	n := len(grid)
	s := make(NeighborSum, n*n)
	for i, row := range grid {
		for j, v := range row {
			for k, d := range dirs {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < n {
					s[v][k/4] += grid[x][y]
				}
			}
		}
	}
	return s
}

func (s NeighborSum) AdjacentSum(value int) int {
	return s[value][0]
}

func (s NeighborSum) DiagonalSum(value int) int {
	return s[value][1]
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
