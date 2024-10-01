package main

import "math/bits"

func maximumRows(matrix [][]int, numSelect int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	mask := make([]int, m)
	for i, row := range matrix {
		for j, x := range row {
			mask[i] |= x << j
		}
	}

	for subset := 0; subset < 1<<n; subset++ {
		if bits.OnesCount(uint(subset)) != numSelect {
			continue
		}
		coveredRows := 0
		for _, row := range mask {
			if row&subset == row {
				coveredRows++
			}
		}
		ans = max(ans, coveredRows)
	}
	return
}

func main() {

}
