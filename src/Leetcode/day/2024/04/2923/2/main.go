package main

/*
	侧成峰
	如果 j 列的元素都是 0，说明没有对于可以击败 j 队
*/

func findChampion(grid [][]int) int {
	for j := range grid[0] {
		ok := true
		for i := 0; i < len(grid) && ok; i++ {
			ok = grid[i][j] == 0
		}
		if ok {
			return j
		}
	}
	return -1
}
