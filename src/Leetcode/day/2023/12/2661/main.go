package main

import "fmt"

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	fmt.Println(m, n)
	idx := make(map[int][2]int)
	for i := range mat {
		for j := range mat[i] {
			idx[mat[i][j]] = [2]int{i, j}
		}
	}

	row, col := make([]int, m), make([]int, n)
	for k := 0; ; k++ {
		x := idx[arr[k]]
		i, j := x[0], x[1]
		row[i]++
		col[j]++
		if row[i] == n || col[j] == m {
			return k
		}
	}
}

func main() {
	fmt.Println(firstCompleteIndex(
		[]int{1, 3, 4, 2},
		[][]int{{1, 4}, {2, 3}},
	))
	fmt.Println(firstCompleteIndex(
		[]int{2, 8, 7, 4, 1, 3, 5, 6, 9},
		[][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}},
	))
	fmt.Println(firstCompleteIndex(
		[]int{1, 4, 5, 2, 6, 3},
		[][]int{{4, 3, 5}, {1, 2, 6}},
	))
}
