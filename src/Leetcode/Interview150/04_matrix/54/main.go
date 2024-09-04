package main

import (
	"math"
)

func spiralOrder(matrix [][]int) (ans []int) {
	n, m := len(matrix), len(matrix[0])
	if n == 0 || m == 0 {
		return nil
	}
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for x, y, k, d := 0, 0, 1, 0; k <= n*m; k++ {
		//fmt.Println(ans, x, y, k, d)
		ans = append(ans, matrix[x][y])
		matrix[x][y] = math.MaxInt
		a, b := x+dx[d], y+dy[d]
		if a < 0 || a >= n || b < 0 || b >= m || matrix[a][b] == math.MaxInt {
			d = (d + 1) % 4
			a, b = x+dx[d], y+dy[d]
		}
		x, y = a, b
	}
	return
}
