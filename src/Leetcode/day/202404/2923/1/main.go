package main

/*
	横看成岭

	如果 grid[i] 有 n-1 个 1，即元素和为 n-1，说明 i 队比其他 n-1 个队都要强，i 队是冠军。
	也可以判断，对于这一行的所有不等于 i 的 j，都有 grid[i][j] = 1。这样可以在遇到 0 的时候，提前退出循环。
*/

func findChampion(grid [][]int) int {
	for i, row := range grid {
		ok := true
		for j := 0; j < len(row) && ok; j++ {
			ok = j == i || row[j] == 1
		}
		if ok {
			return i
		}
	}
	return -1
}
