package main

/*
	打擂台
*/

func findChampion(grid [][]int) (ans int) {
	for i, row := range grid {
		if row[ans] == 1 {
			ans = i
		}
	}
	return
}
