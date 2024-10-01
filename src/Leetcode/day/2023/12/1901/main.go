package main

import (
	"fmt"
	"sort"
)

func findPeakGrid(mat [][]int) []int {
	indexOfMax := func(x []int) int {
		res := 0
		for i := 1; i < len(x); i++ {
			if x[i] > x[res] {
				res = i
			}
		}
		return res
	}
	maxJ := 0
	i := sort.Search(len(mat)-1, func(i int) bool {
		maxJ = indexOfMax(mat[i])
		return mat[i][maxJ] > mat[i+1][maxJ]
	})
	return []int{i, indexOfMax(mat[i])}
}

func main() {
	fmt.Println(findPeakGrid([][]int{{1, 4}, {3, 2}}))
	fmt.Println(findPeakGrid([][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}))
}
